package main

// Coordinates data between the channels
func analyze(input chan *Users, messages chan *Message) {

	arrivals := make(chan *User)   // Delivering users arriving
	departures := make(chan *User) // Delivering users leaving

	messages1 := make(chan *Message)
	messages2 := make(chan *Message)

	go packer(arrivals, "Arrived", messages1)
	go packer(departures, "Departed", messages2)

	fanInMessages(messages1, messages2, messages)

	// Vars holding two user lists
	var users *Users
	var oldUsers *Users

	// Grab a user list to start with
	users = <-input

	for {
		// Shifting and grabbing new user list
		oldUsers = users
		users = <-input

		// Sends differences in user lists to appropriate chanel
		oldUsers.sendUsersMissingFrom(users, arrivals)   // Arriving smurfs
		users.sendUsersMissingFrom(oldUsers, departures) // Leaving smurfs
	}
}

func fanInMessages(input1 chan *Message, input2 chan *Message, output chan *Message) {
	go func() {
		for {
			output <- <-input1
		}
	}()
	go func() {
		for {
			output <- <-input2
		}
	}()
}

func packer(input chan *User, action string, output chan *Message) {
	for {
		output <-&Message{
			action,
			<-input,
		}
	}
}
