package main

// Coordinates data between the channels
func analyze(input chan *Users, arrivals chan *User, departures chan *User) {

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
