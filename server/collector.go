package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)


// Collects user lists and send then back on the chanel
func smurfGetter(url string, smurfRoad chan *Users) {

	const SLEEPTIME time.Duration = time.Second * time.Duration(5)

	for {

		time.Sleep(SLEEPTIME)
		r, e := http.Get(url)
		if e != nil {
			fmt.Println("ERROR(getsmurfs): " + e.Error())
			time.Sleep(time.Second) //Then retry
		} else {

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println("ERROR(getsmurfs): " + e.Error())
				time.Sleep(time.Second) //Then retry
			} else {
				var users Users
				json.Unmarshal(body, &users)

				smurfRoad <- &users
			}

			r.Body.Close()
		}
	}
}
