package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/0xAX/notificator"
	"time"
)

type User struct {
	Id		string	`json:"user_id"`
	Nick		string	`json:"nick"`
}

type Users []User

func (users Users) contains(searchTerm string) (bool)  {
	for _, user := range users {
		if user.Id == searchTerm {
			return true
		}
	}
	return false
}

func main() {


	var notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "StalkIT",
	})

	users := getSmurfs("https://hubbit.chalmers.it/sessions.json")
	for {
		time.Sleep(100000)
		oldUsers := users
		users = getSmurfs("https://hubbit.chalmers.it/sessions.json")


		for _, user := range users {
			if !oldUsers.contains(user.Id) {
				// user arrived
				notify.Push("StalkIT", user.Nick + " has arrived at the Hubb", "/home/user/icon.png", notificator.UR_NORMAL)
			}
		}

		for _, user := range oldUsers {
			if !users.contains(user.Id) {
				// user left
				notify.Push("StalkIT", user.Nick + " has left the Hubb", "/home/user/icon.png", notificator.UR_NORMAL)
			}
		}
	}
}

func getSmurfs(url string) (users Users) {
	r, e := http.Get(url)
	if e != nil{
		fmt.Printf(e.Error())
	}
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &users)

	return users
}