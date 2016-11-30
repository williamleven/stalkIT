package main

import (
	"time"
)

type UserTime struct {
	user *User
	time int
}

type UsersTimes []UserTime

func (users *UsersTimes)add(user User) UsersTimes {
	var newArray UsersTimes
	newArray = make([]UserTime, len(*users) + 1)

	for i, u := range *users {
		newArray[i] = u
	}
	newArray[len(*users)] = UserTime{&user, 0}

	return newArray
}

func (users *UsersTimes)remove(user *User) UsersTimes {
	var newArray UsersTimes
	newArray = make([]UserTime, len(*users) - 1)

	i := 0
	for _, u := range *users {
		if !(u.user.Id == user.Id)  {
			newArray[i] = u
			i++
		}
	}

	return newArray
}

func (users *UsersTimes)contains(searchItem *User) bool {
	for _, u := range *users {
		if u.user.Id == searchItem.Id {
			return true
		}
	}
	return false
}


func departureHandler(missingUsers chan *User, departures chan *User, newUsers chan *User, arrivals chan *User) {
	var toDepart UsersTimes = make(UsersTimes, 0)
	const GRACETIME = 10

	for {
		select {
		case user := <-missingUsers:
			if(!toDepart.contains(user)) {
				toDepart = toDepart.add(*user)
			}

		case user := <-newUsers:
			if(toDepart.contains(user)) {
				toDepart = toDepart.remove(user)
			} else {
				arrivals <-user
			}

		default:
			i := 0
			for i < len(toDepart) {
				if(toDepart[i].time >= GRACETIME) {
					departures <- toDepart[i].user
					toDepart = toDepart.remove(toDepart[i].user)
				} else {
					toDepart[i].time++
					i++
				}
			}

			time.Sleep(time.Second * time.Duration(1))
		}
	}
}
