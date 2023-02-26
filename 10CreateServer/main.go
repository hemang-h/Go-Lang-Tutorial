package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	const url string = "http://[::1]:8000"
	PerformGetRequest(url)
}

func PerformGetRequest(url string) {
	response, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer response.Body.Close()
	var responseBuilder strings.Builder
	dataBytes, _ :=ioutil.ReadAll(response.Body)	
	content, _ := responseBuilder.Write(dataBytes)
	fmt.Println(content)
	fmt.Println()	
}
