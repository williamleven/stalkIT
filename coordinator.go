package main

func coordinate(input chan Users, outputArrivals chan User, outputDepartures chan User)  {

	var users Users
	var oldUsers Users

	users = <-input

	for {
		oldUsers = users
		users = <-input

		go oldUsers.usersMissingFrom(users, outputArrivals) // arriving smurfs
		go users.usersMissingFrom(oldUsers, outputDepartures) // leaving smurfs
	}
}