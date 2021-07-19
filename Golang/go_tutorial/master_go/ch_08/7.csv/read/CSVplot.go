package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a data file!")
		return
	}
	file := os.Args[1]
	_, err := os.Stat(file)
	if err != nil {
		fmt.Println("Cannot stat", file)
		return
	}

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Cannot open", file)
		fmt.Println(err)
		return
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, rec := range allRecords {
		fmt.Printf("x: %s, y: %s\n", rec[0], rec[1])
	}

}
