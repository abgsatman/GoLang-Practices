package main

import (
	"errors"
	"fmt"
)

func traceController() {
	defer fmt.Println("traceController is running....")

	otherTrace()
}
func otherTrace() {
	defer fmt.Println("otherTrace is running...")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Oh, panic: %+v\n", r)
		}
	}()
	anotherTrace()
}
func anotherTrace() {
	defer fmt.Println("anotherTrace is running...")
	panic(errors.New("there is a big boy here"))
}
