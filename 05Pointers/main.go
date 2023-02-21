package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the pointer class")
	// var ptr *int 
	// fmt.Println("The default value is :", ptr) 

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("The value of ptr is: ", ptr)
	fmt.Println("The value of *ptr is: ", *ptr)
}