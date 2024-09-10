package signalcli

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

var onReceiveHandlers = make([]OnReceiveHandler, 0)

type OnReceiveHandler func(payload *Payload)

func OnReceive(handler OnReceiveHandler) {
	onReceiveHandlers = append(onReceiveHandlers, handler)
}

func init() {
	go func() {
		for {
			receive()
			time.Sleep(time.Second * 5)
		}
	}()
}

func receive() {
	cmd := exec.Command("signal-cli", "-o", "json", "receive")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	packets := strings.Split(string(output), "\n")
	for _, packet := range packets {
		if packet == "" {
			continue
		}
		envelope := &Payload{}
		err := json.Unmarshal([]byte(packet), envelope)
		if err != nil {
			fmt.Println(err.Error() + ": " + packet)
			continue
		}
		for _, handler := range onReceiveHandlers {
			handler(envelope)
		}
	}
}
