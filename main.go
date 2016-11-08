package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	//"github.com/0xAX/notificator"
	"time"
)

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

func (users Users) usersMissingFrom(otherUsers Users, output chan User)  {
	for _, user := range otherUsers  {
		if !users.contains(user) {
			output <- user
		}
	}
}

func notifier(input chan User, append string)  {
	for {
		user := <-input
		// user arrived
		fmt.Println(user.Nick + append)
		//notify.Push("StalkIT", user.Nick + " has arrived at the Hubb", "/home/user/icon.png", notificator.UR_NORMAL)
	}
}

func main() {


	/*var notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "StalkIT",
	})*/


	input := make(chan Users)
	outputArrivals := make(chan User)
	outputDepartures := make(chan User)

	var users Users
	var oldUsers Users

	go getSmurfs("https://hubbit.chalmers.it/sessions.json", input)
	users = <-input

	go notifier(outputArrivals, " has arrived at the Hubb")
	go notifier(outputDepartures, " has left the Hubb")

	for {
		go getSmurfs("https://hubbit.chalmers.it/sessions.json", input)

		time.Sleep(time.Second * 10)

		oldUsers = users
		users = <-input

		go oldUsers.usersMissingFrom(users, outputArrivals) // arriving smurfs
		go users.usersMissingFrom(oldUsers, outputDepartures) // leaving smurfs
	}
}

func getSmurfs(url string, smurfRoad chan Users) {
	r, e := http.Get(url)
	if e != nil{
		fmt.Printf(e.Error())
	}
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	var users Users
	json.Unmarshal(body, &users)

	smurfRoad <- users
}