package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/humbertovnavarro/signal-smtp-shim/pkg/mail"
	signalcli "github.com/humbertovnavarro/signal-smtp-shim/pkg/signal-cli"
)

var RecipientMap = make(map[string]string)

func init() {
	usersJson, err := os.ReadFile("./users.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(usersJson, &RecipientMap)
	if err != nil {
		panic(err)
	}
}

func main() {
	mail.OnReceive(func(mail *mail.Mail) {
		if recipient, ok := RecipientMap[mail.To]; ok {
			messageBody := strings.Join(mail.Content, "\n")
			messageHeader := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s", mail.From, mail.To, mail.Subject)
			start := 0
			messageBodyLen := len(messageBody)
			for start < messageBodyLen {
				minIndex := min(messageBodyLen, start+2000)
				chunk := messageBody[start:minIndex]
				err := signalcli.Send(fmt.Sprintf("%s\n%s", messageHeader, chunk), recipient)
				if err != nil {
					fmt.Println(err)
				}
				start = 2000
			}
		}
	})
	signalcli.OnReceive(func(payload *signalcli.Payload) {
		fmt.Println(payload.ToString())
	})
	mail.ListenAndServe("0.0.0.0:25")
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Blocking, press ctrl+c to continue...")
	<-done
}
