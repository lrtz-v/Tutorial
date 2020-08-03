package config

import (
	"context"
	"testing"
)

func TestConnection(t *testing.T) {
	ctx := context.Background()

	esConfig := GetEsInstance(ctx, "books")
	defer esConfig.Stop()

	esConfig.Ping(ctx, "http://127.0.0.1:9200")
}

func TestQueryWithID(t *testing.T) {
	ctx := context.Background()

	esConfig := GetEsInstance(ctx, "books")
	defer esConfig.Stop()

	esConfig.QueryWithID(ctx, "XVSFDBXHHSBJXBNSU")

}

func TestDeleteIndex(t *testing.T) {
	ctx := context.Background()

	esConfig := GetEsInstance(ctx, "books")
	defer esConfig.Stop()
	esConfig.DeleteIndex(ctx)
	
}
