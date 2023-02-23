package main

import "fmt"

var (
	rows int
	columns int
	height int
)


func main() {
	fmt.Println("Hello welcome to the Matrix World")
	fmt.Println("Enter the values for the rows and columns")
	
	fmt.Scan(&rows, &columns, &height)
	
	fmt.Println("Please input the values to the matrix")
    Matrix := inputSlices()

	fmt.Println("The matrix is:")
	outputMatrix(Matrix)

}

func inputSlices() [][][]int{
	Matrix := make([][][]int, rows)
	// fmt.Println("TReached Inside loop")
	for i:= 0; i < rows; i++ {
		// fmt.Println("TReached i loop")
		Matrix[i] = make([][]int, columns)
		// fmt.Println(Matrix[i])
		for j:=0; j < columns; j++ {
			// fmt.Println("TReached J loop")
			Matrix[i][j] = make([]int, height)
			for k:=0; k < height; k++ {
				// fmt.Println("TReached K loop")
				fmt.Scan(&Matrix[i][j][k])
			}
		}
	}
	return Matrix
}

func outputMatrix(Matrix [][][]int) {
	for i:=0; i<rows; i++ {
		for j:=0; j < columns; j++ {
			for k:=0; k < height; k++ {
				fmt.Print(Matrix[i][j][k], " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}