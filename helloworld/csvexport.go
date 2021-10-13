package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func csvExport() {
	legos := [][]string{
		{"lego", "creator"},
		{"lego", "city"},
		{"lego", "technic"},
		{"lego", "idea"},
	}

	file, err := os.Create("src/lego.csv")
	defer file.Close()

	if err != nil {
		fmt.Println("could not open the file", err)
	}

	write := csv.NewWriter(file)
	defer write.Flush()

	for _, lego := range legos {
		if err := write.Write(lego); err != nil {
			fmt.Println("could not write the file", err)
		}
	}
}
