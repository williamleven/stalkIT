package main

// Holds smurf data
type User struct {
	Id   string `json:"user_id"`
	Nick string `json:"nick"`
}

type Users []User

// Checks if the list of users contains a specific user
func (users *Users) contains(searchItem *User) bool {
	for _, user := range *users {
		if user.Id == searchItem.Id {
			return true
		}
	}
	return false
}

// Send all users missing from otherUsers to the channel output
func (users *Users) sendUsersMissingFrom(otherUsers *Users, output chan *User) {
	for _, user := range *otherUsers {
		if !users.contains(&user) {
			output <- &user
		}
	}
}
