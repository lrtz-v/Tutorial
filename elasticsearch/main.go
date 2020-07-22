package main

import (
	"context"
	"elasticsearch/src/book"
	"elasticsearch/src/config"
	"log"
)


func main() {
	ctx := context.Background()
	
	client := config.NewEsClient()
	defer client.Stop()

	for i := 1; i < 2394; i++ {
		b := book.GetBookWithID(ctx, client, i)
		log.Println(b)
	}

}