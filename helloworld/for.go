package main

import (
	"fmt"
)

func forLoop() {
	fmt.Println("for.go file is running............")

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	legoList := make(map[string]string)
	legoList["legocity"] = "Lego City"
	legoList["legocreator"] = "Lego Creator"
	legoList["legoideas"] = "Lego Ideas"

	for index, value := range legoList {
		fmt.Println("index", index, "value: ", value)
	}

	newLegoList := []string{"city", "creator", "ideas", "system", "technic"}

	for index, value := range newLegoList {
		fmt.Println(index, value)
	}
}
