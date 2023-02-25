package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main () {
	defer wg.Wait()
	websites := []string {
		"https://localhost:8080",
		"https://google.com",
		"https://github.com",
		"https://tokenminds.co",
		"https://nethermind.in",
	}
	num := 0
	for _, endpoint := range websites {
		
		wg.Add(1)
		go getStatusCode(endpoint, num)
		// fmt.Printf("Hey number %d!! \n")
		num ++
	}

}

func getStatusCode(endpoint string, num int) {
	defer wg.Done()

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS, something went wrong for: ", endpoint, "and loop number is: ", num)
	} else {
		fmt.Printf("Status code %d for the website %s \n", response.StatusCode, endpoint)
	}
}
