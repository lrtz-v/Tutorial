package book

import (
	"testing"
)

func TestGetBooks(t *testing.T) {
	books := getBooks()
	typeMap := make(map[string]int)

	for _, book := range books {
		t := book.Type
		if _, ok := typeMap[t]; ok {
			typeMap[t]++
		} else {
			typeMap[t] = 1
		}
	}
}