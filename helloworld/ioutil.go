package main

import (
	"fmt"
	"io/ioutil"
)

var url string = "src/ioutil.txt"

func fileOs() {
	data, err := ioutil.ReadFile(url)

	if err != nil {
		fmt.Println("file could not opened")
	}

	ioutil.WriteFile(url, []byte(string(data)+"hello\ngo\n"), 0777)

	fmt.Println(string(data))
}
