package main

import (
	"context"
	"elasticsearch/src/book"
	"elasticsearch/src/config"
)

func main() {
	ctx := context.Background()

	esConfig := config.GetEsInstance(ctx, book.Index)
	defer esConfig.Stop()

}
