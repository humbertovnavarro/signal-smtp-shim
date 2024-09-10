package mail

import (
	"bytes"
	"fmt"
	"log"
	"mime"
	"mime/multipart"
	"mime/quotedprintable"
	"net"
	"net/mail"
	"strings"

	"github.com/mhale/smtpd"
)

type Mail struct {
	Origin  net.Addr
	From    string
	To      string
	Subject string
	Content []string
}

func (m *Mail) String() string {
	return fmt.Sprintf("Origin: %s\nFrom: %s\nTo: %s\nContent:\n%s", m.Origin.String(), m.From, m.To, m.Content)
}

type MailHandlerFunc = func(mail *Mail)

var mailHandlerFuncs = make([]MailHandlerFunc, 0)

func OnReceive(handler MailHandlerFunc) {
	mailHandlerFuncs = append(mailHandlerFuncs, handler)
}

func mailHandler(origin net.Addr, from string, to []string, data []byte) error {
	msg, _ := mail.ReadMessage(bytes.NewReader(data))
	subject := msg.Header.Get("Subject")
	mediaType, params, err := mime.ParseMediaType(msg.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(msg.Body, params["boundary"])
		m := &Mail{
			From:    from,
			To:      to[0],
			Subject: subject,
			Origin:  origin,
		}
		for {
			part, err := mr.NextPart()
			if err != nil {
				break
			}
			defer part.Close()
			encoding := part.Header.Get("Content-Transfer-Encoding")
			var body bytes.Buffer
			if encoding == "quoted-printable" {
				_, err = body.ReadFrom(quotedprintable.NewReader(part))
			} else {
				_, err = body.ReadFrom(part)
			}
			if err != nil {
				log.Fatal(err)
			}
			m.Content = append(m.Content, body.String())
		}
		for _, handler := range mailHandlerFuncs {
			handler(m)
		}
	}
	return nil
}

func ListenAndServe(hostname string) error {
	return smtpd.ListenAndServe("0.0.0.0:25", mailHandler, "SignalSmtpShim", "")
}
