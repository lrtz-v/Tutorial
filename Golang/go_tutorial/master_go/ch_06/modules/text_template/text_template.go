package main

import (
	"fmt"
	"html/template"
	"os"
)

type Entry struct {
	Number int
	Square int
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Need the template file!")
		return
	}

	tFile := args[1]
	DATA := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}

	var Entries []Entry
	for _, i := range DATA {
		if len(i) == 2 {
			temp := Entry{Number: i[0], Square: i[1]}
			Entries = append(Entries, temp)
		}
	}
	t := template.Must(template.ParseGlob(tFile))
	t.Execute(os.Stdout, Entries)
}
