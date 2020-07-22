package config

import (
	elastic "github.com/olivere/elastic/v7"
)

// NewEsClient create new connection to es node
func NewEsClient() *elastic.Client {
	client, _ := elastic.NewClient()
	return client
}
