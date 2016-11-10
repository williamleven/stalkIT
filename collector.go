package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"time"
)

func collector(delay int, url string, output chan Users){
	go getSmurfs(url, output)
	for {
		time.Sleep(time.Second * time.Duration(delay))
		go getSmurfs(url, output)
	}
}

func getSmurfs(url string, smurfRoad chan Users) {
	r, e := http.Get(url)
	if e != nil{
		fmt.Printf("ERROR(getsmurfs): " + e.Error())
		return
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("ERROR(getsmurfs): " + e.Error())
		return
	}
	var users Users
	json.Unmarshal(body, &users)

	smurfRoad <- users
}