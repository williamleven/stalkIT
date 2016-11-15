package main

import (
	"fmt"
	"net"
	"bufio"
	"github.com/0xAX/notificator"
)

func notifier() {
	conn, err := net.Dial("tcp", "localhost:7825")

	if(err != nil) {
		return;
	}

	arrivalMessage := " has arrived at the Hubb"
	departureMessage := " has left the Hubb"

	// Build notificator
	var notify = notificator.New(notificator.Options{
		AppName:     "StalkIT",
	})

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			return;
		}

		message := data[1:(len(data)-1)]
		if data[0] == 'A' {
			message += arrivalMessage
		} else {
			message += departureMessage
		}
		fmt.Println(message)
		notify.Push("StalkIT", message, "icon/default.svg", notificator.UR_NORMAL)
	}
}
