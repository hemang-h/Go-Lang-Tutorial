package main

import "fmt"

func main() {
	fmt.Println("A")
	defer fmt.Println("B")
	fmt.Println("C")
	defer fmt.Println("D")
	fmt.Println("E")
	fmt.Println("F")
    defer fmt.Println("G")
	defer fmt.Println("H")
	fmt.Println("I")
	fmt.Println("J")
	fmt.Println("K")
	defer fmt.Println("L")
}

// The output will be 

/*
	A C E F I J K L H G D B
*/