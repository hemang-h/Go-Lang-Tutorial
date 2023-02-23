package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	const url string = "https://wakatime.com/dashboard"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("type of response is: %T \n", response)
	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataBytes)
	fmt.Println(content)
}