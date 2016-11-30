package main

import (
	"fmt"
	"net"
	"bufio"
	"encoding/json"
	"github.com/0xAX/notificator"
	"github.com/spacemonkeygo/openssl"
)

func notifier() {
	ctx, err := NewCtx()
	if err != nil {
		panic(err)
	}
	err = ctx.LoadVerifyLocations("/etc/ssl/certs/ca-certificates.crt", "")
	if err != nil {
		panic(err)
	}
	conn, err := openssl.Dial("tcp", "stalkit.gurgy.me:4242", ctx, 0)

	if(err != nil) {
		panic(err)
	}

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

		var output string
		if message.Action == "Arrived" {
			output = fmt.Sprintf(language.getPhrase("arrival_message"), message.User.Nick)
		} else {
			output = fmt.Sprintf(language.getPhrase("departure_message"), message.User.Nick)
		}

		fmt.Println(output)
		notify.Push("StalkIT", output, "icon/default.svg", notificator.UR_NORMAL)
	}
}
