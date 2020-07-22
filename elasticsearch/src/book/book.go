package book

import (
	"io/ioutil"
	"strings"
	"time"
)

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"book":{
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
					"type":"geo_point"
				}
			}
		}
	}
}`

// DIR of files
const DIR = ""

type Book struct {
    ID int64 `json:"id"`
    Name string `json:"name"`
    Type string `json:"type"`
    Size int64 `json:"size"`
    Created time.Time `json:"Created"`
}


func getBooks() []Book {
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
