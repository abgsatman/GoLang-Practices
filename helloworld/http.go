package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpServer() {
	fmt.Println("checking...")

	url := "http://www.phpservisi.com/golang/content.txt"

	data := getHTTPResponse(url)

	fmt.Println(data)
}

func getHTTPResponse(url string) string {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("an error occurred while getting response!")
	}

	defer response.Body.Close()

	r, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("an error occurred while getting content!")
	}

	return string(r)
}
