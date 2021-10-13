package main

import (
	"fmt"
	"time"
)

func channels() {
	dataChannel := make(chan string)
	dataChannel2 := make(chan string)

	go sender(dataChannel, "lego creator", time.Millisecond, 250)
	go sender(dataChannel2, "lego city", time.Second, 2)

	for {
		select {
		case message := <-dataChannel:
			fmt.Println(message)
		case message2 := <-dataChannel2:
			fmt.Println(message2)
		}
	}
}

func sender(channel chan<- string, msg string, timeType time.Duration, timeValue float32) {
	for {
		channel <- msg
		time.Sleep(timeType * time.Duration(timeValue))
	}
}
