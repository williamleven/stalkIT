package main

type User struct {
	Id		string	`json:"user_id"`
	Nick		string	`json:"nick"`
}

type Users []User

func (users Users) contains(searchItem User) (bool)  {
	for _, user := range users {
		if user.Id == searchItem.Id {
			return true
		}
	}
	return false
}

func (users Users) usersMissingFrom(otherUsers *Users, output chan *User)  {
	for _, user := range *otherUsers  {
		if !users.contains(user) {
			output <- &user
		}
	}
}