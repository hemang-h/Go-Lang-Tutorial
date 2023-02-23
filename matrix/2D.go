package main

import "fmt"

// var matrix = [][]int{}

func main() {
	fmt.Println("Hello, let's print a matrix")
	fmt.Println("Enter the values for the row and columns")
	rows, columns := values()
	
	fmt.Println("Now please enter the values that you want to see in the matrix") 
	matrix := inputMatrix(rows, columns)	
	outputMatrix(rows, columns, matrix)
}

func values () (int, int) {
	var rows int
	var columns int

	fmt.Scan(&rows, &columns)
	return rows, columns
}

func inputMatrix(rows int, columns int) [][]int{
	// var values int
	matrix := make([][]int, rows)	
	for i := 0; i < rows; i ++ {
		matrix[i] = make([]int, columns )
		fmt.Println("matrix[i] is:", matrix[i])
		for j:=0; j < columns; j++ {
			fmt.Scan(&matrix[i][j])			
		}
	}
	return matrix
}
func outputMatrix(rows int, columns int, matrix [][]int) {
	for i:=0; i < rows; i++ {
		for j:=0; j< columns; j++{
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}