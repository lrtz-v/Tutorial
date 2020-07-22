package book

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	elastic "github.com/olivere/elastic/v7"
)

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"books":{
			"properties":{
                "id": {
                    "type": "long"
                },
				"name":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"type":{
					"type":"keyword"
				},
				"size":{
					"type":"long"
				}
			}
		}
	}
}`

// DIR of files
const DIR = "./mybooks"

// Book define struct of books
type Book struct {
    ID int64 `json:"id"`
    Name string `json:"name"`
    Type string `json:"type"`
    Size int64 `json:"size"`
    Created time.Time `json:"Created"`
}

// GetBooks return all books in $DIR
func GetBooks() []Book {
	files, err := ioutil.ReadDir(DIR)
    if err != nil {
        panic(err)
    }

    books := make([]Book, len(files))

    for i := 0; i < len(files); i++ {
        t := strings.Split(files[i].Name(), ".")
        books[i] = Book{
            ID: int64(i),
            Name: strings.Join(t[0:len(t)-1], "."), 
            Type: t[len(t)-1],
            Size: files[i].Size(), 
            Created: files[i].ModTime(),
        }
    }

    return books
}

// UploadBooks upload books to es
func UploadBooks(ctx context.Context, client *elastic.Client) {
	// indexCheck(ctx, client)
	books := GetBooks()

	for _, book := range books {
		info, err := client.Index().
			Index("books").
			Id(strconv.Itoa(int(book.ID))).
			BodyJson(book).
			Do(ctx)
		if err != nil {
			log.Fatalf("Indexed Failed, %d", book.ID)
			panic(err)
		}
		log.Printf("Indexed books %s to index %s\n", info.Id, info.Index)
	}
}

// GetBookWithID query book by bookId
func GetBookWithID(ctx context.Context, client *elastic.Client, id int) Book {
	res, err := client.Get().
		Index("books").
		Id(strconv.Itoa(id)).
		Do(ctx)
	if err != nil {
		log.Fatalf("[*]GetBookWithID Failed, %d", id)
		panic(err)
	}
	if res.Found {
		log.Printf("Got document %s in version %d from index %s\n", res.Id, res.Version, res.Index)
	}

	var book Book

	err = json.Unmarshal(res.Source, &book)
	if err != nil {
		log.Fatalln("[*]Source Unmarshal Failed")
		panic(err)
	}
	return book
}