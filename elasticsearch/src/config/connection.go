package config

import (
	elastic "github.com/olivere/elastic/v7"
)

// NewEsClient create new connection to es node
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
