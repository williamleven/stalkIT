package main

import (
	"fmt"
	"net"
	"bufio"
	"encoding/json"
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
		data, err := bufio.NewReader(conn).ReadString('\n') // TODO: Parse json-object whilst receiving data to better determine end-of-message

		if err != nil {
			return;
		}

		var message Message
		// Parse json-object, ignore final '\n'-character
		json.Unmarshal([]byte(data[:(len(data)-1)]), &message)

		output := message.User.Nick
		if message.Action == "Arrived" {
			output += arrivalMessage
		} else {
			output += departureMessage
		}

		fmt.Println(output)
		notify.Push("StalkIT", output, "icon/default.svg", notificator.UR_NORMAL)
	}
}
