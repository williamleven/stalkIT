package main

import (
	"github.com/0xAX/notificator"
)

func notifier(input chan *User, append string)  {

	var notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "StalkIT",
	})

	for {
		user := <-input
		notify.Push("StalkIT", user.Nick + append, "/home/user/icon.png", notificator.UR_NORMAL)
	}
}