package main

import (
	"fmt"
)

func print() {
	fmt.Println("printing with a line")

	i := 2
	s := "golang"
	f := 3.14

	fmt.Printf("integer: %d\n", i)
	fmt.Printf("string: %s\n", s)
	fmt.Printf("float: %f\n", f)

	fmt.Println("------------------------------------------")

	fmt.Printf("the type of the first one is: %[1]T\n", i)
	fmt.Printf("the type of the second one is: %[1]T\n", s)
	fmt.Printf("the type of the last one is: %[1]T\n", f)

	fmt.Println("------------------------------------------")

	fmt.Printf("the value of the first one is: %-4v\n", i)
	fmt.Printf("the value of the second one is: %-4v\n", s)
	fmt.Printf("the value of the last one is: %-4v\n", f)
}
