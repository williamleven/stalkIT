package main

import (
	"github.com/0xAX/notificator"
)

// Sends notification in the format input + append
func notifier(input chan *User, append string)  {

	// Creating notificator
	var notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "StalkIT",
	})

	// Pushing notifications whenever receiving input
	for {
		user := <-input
		notify.Push("StalkIT", user.Nick + append, "/home/user/icon.png", notificator.UR_NORMAL)
	}
}