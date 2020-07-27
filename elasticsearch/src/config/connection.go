package config

import (
	"context"
	"log"

	elastic "github.com/olivere/elastic/v7"
)

// EsClient is defintion of es connection management
type EsClient struct {
	Client *elastic.Client
	Index string
}

// NewEsClient return new connection to es
func NewEsClient() *elastic.Client {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	return client
}

// GetEsInstance return instance of connection to es
func GetEsInstance(index string) *EsClient {
	client := EsClient{Index: index}
	client.Client = NewEsClient()
	return &client
}

// Stop connection to es
func (e *EsClient) Stop() {
	e.Client.Stop()
}

// QueryWithID query documents with ID of documents
func (e *EsClient) QueryWithID(ctx context.Context, id string) (*elastic.GetResult, error) {
	res, err := e.Client.Get().
		Index(e.Index).
		Id(id).
		Do(ctx)
	return res, err
}

// Query use DSL to query
func (e *EsClient) Query(ctx context.Context, query *elastic.BoolQuery) ([]*elastic.SearchHit) {
	res, err := e.Client.Search().Index(e.Index).Query(query).Do(ctx)
	if err != nil {
		panic(err)
	}
	return res.Hits.Hits
}

// Indexs create document
func (e *EsClient) Indexs(ctx context.Context, ID string, body interface{}) {
	info, err := e.Client.Index().Index(e.Index).BodyJson(body).Do(ctx)
	if err != nil {
		log.Fatalf("Indexs Failed!, ID: %s\n", ID)
		panic(err)
	}
	log.Printf("Indexed books %s to index %s\n", info.Id, info.Index)
}
