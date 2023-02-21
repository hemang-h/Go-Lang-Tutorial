package main

import "fmt"

type Friend struct {
	friendName string
	age uint
	maritalStatus bool
}

var friendName string
var age uint
var maritalStatus bool
var friend = make([]Friend, 0) 

func main() {
	for{
		fmt.Println("Enter your friends details")
		friendDetail()	
		displayDetails()
	}
}


func friendDetail() {
	fmt.Scan(&friendName, &age, &maritalStatus)
}

func displayDetails() {
	var friends = Friend {
		friendName: friendName,
		age: age,
		maritalStatus: maritalStatus,
	}

	friend = append(friend, friends)
	fmt.Println(friend)
}