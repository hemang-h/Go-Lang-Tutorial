// COMMA OK SYNTAX


package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Welcome to Pizza Review")
	fmt.Println("Enter the ratings for the Pizza")

	input, _ := reader.ReadString('\n')
	fmt.Printf("Thanks for giving us a %v ratings \n", input)
}