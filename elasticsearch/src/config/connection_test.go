package config

import (
	"context"
	"testing"

	"github.com/olivere/elastic/v7"
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

func TestExplain(t *testing.T) {
	ctx := context.Background()
	esConfig := GetEsInstance(ctx, "books")
	defer esConfig.Stop()
	
	query := elastic.NewBoolQuery()
	query.Should(elastic.NewMatchQuery("name", "平等"))
	hits := esConfig.Explain(ctx, query)
	t.Fatal(hits[0].Explanation)
}
