package main

import (
	"fmt"
	"time"
)

func goroutine() {
	fmt.Println("checking....")

	go worker("lego creator")
	go worker("lego city")
	worker("lego art")
}

func worker(word string) {
	if word == "" {
		fmt.Println("string cannot be empty!")
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 250)
		fmt.Println(word)
	}
}
