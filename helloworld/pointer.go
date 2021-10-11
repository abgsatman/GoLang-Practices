package main

import "fmt"

func pointer() {
	var a int = 3

	fmt.Println("a is just getting started as:", a) //printing a as default

	d := inc(&a)
	fmt.Println("a is changed as +1 with pointer use:", d, "but a is now:", a) //printing a as incremental with pointer use

	c := otherInc(a)
	fmt.Println("a is changed as +1 by default", c, "but a is now:", a) //printing a as incremental by default

	b := otherFunc(a)
	fmt.Println("a is the result as:", b, "but a is now:", a) //printing a by default

	e := otherInc(a)
	fmt.Println("a is changed as +1 by default", e, "but a is now:", a) //printing a as incremental by default
}

func otherInc(param int) int {
	param++
	return param
}

func inc(i *int) int {
	*i++
	return *i
}

func otherFunc(param int) int {
	return param
}
