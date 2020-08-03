package book

import (
	"context"
	"elasticsearch/src/config"
	"testing"
	"time"
)

func TestGetBooks(t *testing.T) {
	books := GetBooks()
	typeMap := make(map[string]int)

	for _, book := range books {
		t := book.Type
		if _, ok := typeMap[t]; ok {
			typeMap[t]++
		} else {
			typeMap[t] = 1
		}
	}
}

func TestMarshal(t *testing.T) {
	book := Book{ID: int64(1), Name: "Book", Type: "pdf", Size: 1024, Created: time.Now()}

	str := Marshal(book)
	if str == nil {
		t.Fatal("Test Marshal book Failed.")
	}

	book2 := Unmarshal(str)
	if book2 == nil {
		t.Fatal("Test Unmarshal book Failed.")
	}

	if book2.ID != book.ID || book2.Name != book.Name || book.Size != book2.Size || book.Created != book2.Created {
		t.Fatal("Test Unmarshal err.")
	}
}

func TestGetBookWithID(t *testing.T) {
	ctx := context.Background()

	esConfig := config.GetEsInstance(ctx, Index)
	defer esConfig.Stop()

	book := GetBookWithID(ctx, esConfig, "E79ktHMB--pMjvYfOjDB")
	if book == nil {
		t.Fatal("Test GetBookWithID Error.")
	}
}

func TestGetBookWithBookID(t *testing.T) {
	ctx := context.Background()

	esConfig := config.GetEsInstance(ctx, Index)
	defer esConfig.Stop()

	book := GetBookWithBookID(ctx, esConfig, 1)
	if book == nil {
		t.Fatal("TestGetBookWithBookID Error.")
	}
}

func TestGetBookWithName(t *testing.T) {
	ctx := context.Background()

	esConfig := config.GetEsInstance(ctx, Index)
	defer esConfig.Stop()

	books := GetBookWithName(ctx, esConfig, "1984")
	if books == nil || len(books) == 0 {
		t.Fatal("Test GetBookWithName Error.")
	}
}
