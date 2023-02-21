// COMMA OK SYNTAX

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Welcome to Pizza Review")
	fmt.Println("Enter the ratings for the Pizza between 1 & 5")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for giving us ratings ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating)
		numRating ++
		fmt.Println(numRating) 
	}	 
}