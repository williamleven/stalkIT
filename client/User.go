package main

import "time"

// Holds smurf data
type User struct {
	Id   		string		`json:"user_id"`
	Nick 		string 		`json:"nick"`
	StartTime	time.Time	`json:"start_time"`
}
