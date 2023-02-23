package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to the file creating program")
	content := "This is going to go in my file"

	file, err := os.Create("./MyFile.txt")
	if (err != nil) {
		panic(err)
	}

	fmt.Printf("The type is: %T \n", file)

	length, err := io.WriteString(file, content)
	if (err != nil) {
		panic(err)
	}
	fmt.Println("Length is:", length)
	defer file.Close()

	readFile("./MyFile.txt")

}

func readFile (filename string) {

	dataByte, err := ioutil.ReadFile(filename)
	if (err != nil) {
		panic(err)
	}
	fmt.Println("Filename is:", string(dataByte))
}