package main

import (
	"fmt"
)

func uploadBooks() {
    // books := getBooks()
    // client := config.NewEsClient()
}


func main() {
	books := book.getBooks()	
	for _, book := range books {
		t := book.Type
		if _, ok := typeMap[t]; ok {
			typeMap[t]++
		} else {
			typeMap[t] = 1
		}
	}
	fmt.Printf(typeMap)
}