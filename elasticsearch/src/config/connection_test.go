package config

import (
	"context"
	"testing"

	"elasticsearch/src/book"
)

func TestConnection(t *testing.T) {
	ctx := context.Background()

	esConfig := GetEsInstance(ctx, book.Index)
	defer esConfig.Stop()

	esConfig.Ping("http://127.0.0.1:9200")
}

func TestQueryWithID(t *testing.T) {
	ctx := context.Background()

	esConfig := GetEsInstance(ctx, book.Index)
	defer esConfig.Stop()

	QueryWithID(ctx, "XVSFDBXHHSBJXBNSU")

}
