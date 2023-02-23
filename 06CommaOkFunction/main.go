package main

import "fmt"

func main() {
	result, myMessage := continuousAddition(2, 4, 5, 6)
	fmt.Println("Result is: ", result)
	fmt.Println(myMessage)
}

func continuousAddition(values ...int) (int, string) {
	total := 0
	for _, sum := range values {
		total = total + sum
	}
	return total, "Catch me if you can"
}

