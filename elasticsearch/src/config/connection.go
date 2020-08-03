package config

import (
	"context"
	"log"

	elastic "github.com/olivere/elastic/v7"
)

// EsClient is defintion of es connection management
type EsClient struct {
	Client *elastic.Client
	Index  string
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

// Ping specify url
func (e *EsClient) Ping(ctx context.Context, url string) (*elastic.PingResult, int) {
	info, code, err := e.Client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		log.Fatalf("Ping [%s] Error!", url)
		panic(err)
	}
	return info, code
}

// GetEsInstance return instance of connection to es
func GetEsInstance(ctx context.Context, index string) *EsClient {
	client := EsClient{Index: index}
	client.Client = NewEsClient()
	client.CreateIndex(ctx)
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
func (e *EsClient) Query(ctx context.Context, query *elastic.BoolQuery) []*elastic.SearchHit {
	res, err := e.Client.Search().Index(e.Index).Query(query).Do(ctx)
	if err != nil {
		panic(err)
	}
	return res.Hits.Hits
}

// Indexs create document
func (e *EsClient) Indexs(ctx context.Context, body interface{}) {
	info, err := e.Client.Index().Index(e.Index).BodyJson(body).Do(ctx)
	if err != nil {
		log.Fatalf("Indexs Failed!, body: %s\n", body)
		panic(err)
	}
	log.Printf("Indexed books %s to index %s\n", info.Id, info.Index)
}

// BulkInsert multi insert
func (e *EsClient) BulkInsert(ctx context.Context, body []interface{}) {
	bulkRequest := e.Client.Bulk()
	for _, v := range body {
		req := elastic.NewBulkIndexRequest().Index(e.Index).Doc(v)
		bulkRequest = bulkRequest.Add(req)
	}
	_, err := bulkRequest.Do(ctx)
	if err != nil {
		log.Printf("BulkInsert Failed, err: %s", err)
		panic(err)
	}
}

// CreateIndex create index if not exist
func (e *EsClient) CreateIndex(ctx context.Context) bool {
	exist, err := e.Client.IndexExists(e.Index).Do(ctx)
	if err != nil {
		log.Fatalf("Check Index exists failed, err: %s", err)
		panic(err)
	}
	if exist {
		return true
	}
	res, err := e.Client.CreateIndex(e.Index).Do(ctx)
	if err != nil {
		log.Fatalf("Create Index failed, err: %s", err)
		panic(err)
	}
	if !res.Acknowledged {
		log.Println("Create Index failed")
		return false
	}
	return true
}

// CreateMapping add mapping
func (e *EsClient) CreateMapping(ctx context.Context, mapping string) {
	res, err := e.Client.PutMapping().Index(e.Index).BodyString(mapping).Do(ctx)
	if err != nil {
		log.Printf("CreateMapping Failed: %s", err)
		panic(err)
	}

	if !res.Acknowledged {
		log.Println("CreateMapping Failed")
		panic("CreateMapping Failed")
	}
}

// GetMapping return mapping of index
func (e *EsClient) GetMapping(ctx context.Context) map[string]interface{} {
	res, err := e.Client.GetMapping().Index(e.Index).Do(ctx)
	if err != nil {
		log.Printf("GetMapping Failed: %s", err)
		panic(err)
	}

	return res
}

// DeleteIndex delete index
func (e *EsClient) DeleteIndex(ctx context.Context) {
	_, err := e.Client.DeleteIndex(e.Index).Do(ctx)
	if err != nil {
		log.Fatalf("Delete Index [%s] Failed.", e.Index)
		panic(err)
	}
}

// Explain query dsl
func (e *EsClient) Explain(ctx context.Context, query *elastic.BoolQuery) []*elastic.SearchHit {
	res, err := e.Client.Search().Index(e.Index).Query(query).Explain(true).Do(ctx)
	if err != nil {
		panic(err)
	}
	return res.Hits.Hits
}
