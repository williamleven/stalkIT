package main

import (
	"time"
	"strconv"
	"math/rand"
)

// Holds smurf data
type User struct {
	Id   		string		`json:"user_id"`
	Nick 		string 		`json:"nick"`
	StartTime	time.Time	`json:"start_time"`
}

// Checks if the list of users contains a specific user
func contains(users []*User, searchItem string) bool {
	for _, user := range users {
		if user.Id == searchItem {
			return true
		}
	}
	return false
}

func initiate(n int) (users []*User) {
	var user *User
	for i := 0; i <= n ;i++  {
		user = &User{
			Id: strconv.Itoa(i),
			Nick: "some" + strconv.Itoa(i) + "Nick",
			StartTime: time.Now(),
		}
		users = append(users, user)
	}
	return users
}

func removeRandom(users []*User) ([]*User) {
	for n := rand.Intn(3); n > 0 ;n--  {
		if len(users) > 0 {

			i := rand.Intn(len(users)-1)

			users[i] = users[len(users)-1]
			users[len(users)-1] = nil
			users = users[:len(users)-1]
		}
	}
	return users
}

func addRandom(users []*User) ([]*User) {
	for n := rand.Intn(3); n > 0 ;n--  {

		i := rand.Intn(len(users) * 100 + 1)
		for contains(users, strconv.Itoa(i))  {
			i = rand.Intn(len(users)-1)
		}

		user := &User{
			Id: strconv.Itoa(i),
			Nick: "some" + strconv.Itoa(i) + "Nick",
			StartTime: time.Now(),
		}
		users = append(users, user)
	}
	return users
}
