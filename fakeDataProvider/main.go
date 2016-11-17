package main

import (
	"net/http"
	"encoding/json"
)
var users []*User

func handler(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(users)
	w.Write(data);
	users = removeRandom(users)
	users = addRandom(users)
}

func main() {
	users = initiate(100)
	http.HandleFunc("/data.json", handler)
	http.ListenAndServe(":8080", nil)
}