package config

import (
	"context"
	"log"
	"testing"
)

func TestConnection(t *testing.T) {

	ctx := context.Background()
	
	client := NewEsClient()
	defer client.Stop()

	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		t.Fatal("Ping Error!")
	}

	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		t.Fatal("Version check Error!")
	}
	log.Printf("Elasticsearch version %s\n", esversion)
}