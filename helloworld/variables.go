package main

import (
	"fmt"
)

func vars() {
	var i int = 12
	j := 13
	var k int
	fmt.Printf("%d %d %d\n", i, j, k)

	var a string = "lego"
	var b string //not nil
	fmt.Printf("%s %s\n", a, b)

	if b == "" {
		fmt.Println("b is empty")
	}

	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", a, a)

	var t bool
	fmt.Printf("%T %v\n", t, t)

	var y float64
	fmt.Printf("%T %v\n", y, y)

}
