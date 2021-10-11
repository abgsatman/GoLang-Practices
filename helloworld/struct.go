package main

import "fmt"

type lego struct {
	legoType string
	minAge   int
}

func structRunner() {
	fmt.Println("checking.....")

	newLego := legoRunner()
	fmt.Println(newLego.legoType, newLego.minAge)

	newestLego := legoController(&newLego)
	fmt.Println(newestLego.legoType, newestLego.minAge)

	fmt.Println(newLego)
}

func legoRunner() lego {
	return lego{"creator", 6}
}

func legoController(n *lego) lego {
	n.legoType = "city"
	return *n
}
