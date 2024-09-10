package signalcli

import (
	"io"
	"os/exec"
	"strings"
)

func Send(message string, recipient string) error {
	isGroup := strings.HasSuffix(recipient, "=")
	isPhone := strings.HasPrefix(recipient, "+")
	isUser := !isGroup && !isPhone
	args := make([]string, 0)
	args = append(args, "send")
	args = append(args, "--message-from-stdin")
	if isGroup {
		args = append(args, "-g")
	} else if isUser {
		args = append(args, "-u")
	}
	args = append(args, recipient)
	cmd := exec.Command("signal-cli", args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}
	_, err = io.WriteString(stdin, message)
	if err != nil {
		return err
	}
	return stdin.Close()
}
