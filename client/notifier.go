package main

import (
	"fmt"
	"bufio"
	"io/ioutil"
	"encoding/json"
	"github.com/0xAX/notificator"
	"crypto/tls"
	"crypto/x509"
)

func notifier() {
	rootPEM, err := ioutil.ReadFile("public.cer")
 	if err != nil {
		panic("Failed to read certificate file")
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("Failed to parse root certificate")
	}

	conn, err := tls.Dial("tcp", "stalkit.gurgy.me:7825", &tls.Config{
		RootCAs: roots,
	})
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

		if message.User == nil {
			continue
		}

		output := message.User.Nick
		if message.Action == "Arrived" {
			output = fmt.Sprintf(language.getPhrase("arrival_message"), message.User.Nick)
		} else {
			output = fmt.Sprintf(language.getPhrase("departure_message"), message.User.Nick)
		}

		fmt.Println(output)
		notify.Push("StalkIT", output, "icon/default.svg", notificator.UR_NORMAL)
	}
}
