package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"time"
)

// Allowing a maximum of one collection to occur / delay
func collectorTimer(delay int, url string, output chan *Users){
	v := make(chan *Users)
	go smurfGetter(&url, v)
	for {
		time.Sleep(time.Second * time.Duration(delay))
		output <- <-v
	}
}

// Collects user lists and send then back on the chanel
func smurfGetter(url *string, smurfRoad chan *Users) {
	for {
		r, e := http.Get(*url)
		if e != nil{
			fmt.Println("ERROR(getsmurfs): " + e.Error())
			time.Sleep(time.Second)
		}else{
			defer r.Body.Close()

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println("ERROR(getsmurfs): " + e.Error())
				time.Sleep(time.Second)
			}else{
				var users Users
				json.Unmarshal(body, &users)

				smurfRoad <- &users
			}
		}
	}
}