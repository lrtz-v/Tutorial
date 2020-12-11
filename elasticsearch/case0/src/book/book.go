package book

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"elasticsearch/src/config"

	elastic "github.com/olivere/elastic/v7"
)

// DIR of files
const (
	JSONFileDir = "./books.json"
	DIR = "./mybooks"
	Index = "books"
)

// Book define struct of books
type Book struct {
    ID int64 `json:"id"`
    Name string `json:"name"`
    Type string `json:"type"`
    Size int64 `json:"size"`
    Created time.Time `json:"created"`
}

// Unmarshal decode source
func Unmarshal(source json.RawMessage) *Book {
	var book Book
	err := json.Unmarshal(source, &book)
	if err != nil {
		log.Fatalln("[*]Source Unmarshal Failed")
		panic(err)
	}
	return &book
}

// Marshal encode source
func Marshal(book Book) []byte {
	bytes, err := json.Marshal(book)
	if err != nil {
		log.Fatalln("[*]Book Marshal Failed")
		panic(err)
	}
	return bytes
}

// GetBooks return all books in $DIR
func GetBooks() []Book {
	jsonFile, err := os.Open(JSONFileDir)
	defer jsonFile.Close()
	if err != nil {
		log.Fatalln("[*]TestPaserJsonFile Failed")
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	books := []Book{}
	json.Unmarshal(byteValue, &books)
	return books
}

// UploadBooks upload books to es
func UploadBooks(ctx context.Context, esConfig *config.EsClient) {
	books := GetBooks()
	for _, book := range books {
		esConfig.Indexs(ctx, book)
	}
}

// GetBookWithID query book by _id
func GetBookWithID(ctx context.Context, esConfig *config.EsClient, id string) *Book {
	res, err := esConfig.QueryWithID(ctx, id)
	if err != nil {
		log.Fatalf("[*]GetBookWithID Failed, _id: %s\n", id)
		panic(err)
	}
	if !res.Found {
		log.Fatalf("[*]GetBookWithID Found Empty, _id: %s\n", id)
	}

	return Unmarshal(res.Source)
}

// GetBookWithBookID query book by bookId
func GetBookWithBookID(ctx context.Context, esConfig *config.EsClient, id int) *Book {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("id", strconv.Itoa(id)))
	
	hits := esConfig.Query(ctx, query)
	if len(hits) == 0 {
		return nil
	}
	return Unmarshal(hits[0].Source)
}

// QueryBookWithName query books with book name
func QueryBookWithName(ctx context.Context, esConfig *config.EsClient, name string) []*Book {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("name", name))

	hits := esConfig.Query(ctx, query)
	books := make([]*Book, len(hits))
	for i, v := range hits {
		books[i] = Unmarshal(v.Source)
	}
	return books
}

// QueryBookWithTermName get books with term book name
func QueryBookWithTermName(ctx context.Context, esConfig *config.EsClient, name string) []*Book {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("name.raw", name))

	hits := esConfig.Query(ctx, query)
	books := make([]*Book, len(hits))
	for i, v := range hits {
		books[i] = Unmarshal(v.Source)
	}
	return books
}