package main

import (
	"bufio"
	"fmt"
	"os"
)

func scanner() {
	fmt.Print("Enter command: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		if scanner.Text() == "q" {
			os.Exit(1)
		} else if scanner.Text() == "help" {
			fmt.Println("Type 'q' to quit!")
		} else {
			fmt.Println("Your command: ", scanner.Text())
		}
		fmt.Print("Enter text: ")
	}

	if scanner.Err() != nil {
		fmt.Println("An Error occurred")
	}
}
