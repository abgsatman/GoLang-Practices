package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func csvOs() {
	csvExport()
	csvImport()
}

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

func csvImport() {
	file, err := os.Open("src/lego.csv")

	if err != nil {
		fmt.Println("could not open the file")
	}

	r := csv.NewReader(file)

	record, err := r.ReadAll()
	if err != nil {
		fmt.Println("could not read the file")
	}

	for i := 0; i < len(record); i++ {
		fmt.Println(record[i])
	}

	//fmt.Println(record[0][1]) value of row 1, column 1
}
