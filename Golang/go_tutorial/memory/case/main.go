package main

import (
	"debug/elf"
	"log"
)

func main() {
	f, err := elf.Open("main")

	if err != nil {
		log.Fatal(err)
	}

	for _, section := range f.Sections {
		log.Println(section)
	}
}
