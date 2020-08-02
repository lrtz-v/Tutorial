package main

import (
	"context"
	"elasticsearch/src/book"
	"elasticsearch/src/config"
	"log"
)


func main() {
	ctx := context.Background()

	esConfig := config.GetEsInstance(ctx, book.Index)
	defer esConfig.Stop()

	book.UploadBooks(ctx, esConfig)

	// b := book.GetBookWithID(ctx, esConfig, 100)
	// log.Println(b.Name)

	books := book.GetBookWithName(ctx, esConfig, "忏悔录")
	for _, v := range books {
		log.Println(v.Name)
	}

}