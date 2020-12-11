package book

import (
	"context"
	"elasticsearch/src/config"
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

func TestGetBooks(t *testing.T) {
	books := GetBooks()
	if len(books) == 0 {
		t.Fatal("[*] Gor Empty Book List.")
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

	if book2.ID != book.ID || book2.Name != book.Name || book.Size != book2.Size {
		t.Fatal("Test Unmarshal err.")
	}
}

func TestGetBookWithID(t *testing.T) {
	ctx := context.Background()

	esConfig := config.GetEsInstance(ctx, Index)
	defer esConfig.Stop()

	book := GetBookWithID(ctx, esConfig, "kgMn43MB_0UsgHt0t7Ja")
	if book == nil {
		t.Fatal("Test GetBookWithID Error.")
	}
}

func TestGetBookWithBookID(t *testing.T) {
	ctx := context.Background()

	esConfig := config.GetEsInstance(ctx, Index)
	defer esConfig.Stop()

	book := GetBookWithBookID(ctx, esConfig, 146)
	if book == nil {
		t.Fatal("TestGetBookWithBookID Error.")
	}
}

func TestGetBookWithName(t *testing.T) {
	ctx := context.Background()

	esConfig := config.GetEsInstance(ctx, Index)
	defer esConfig.Stop()

	books := QueryBookWithName(ctx, esConfig, "1984")
	if books == nil || len(books) == 0 {
		t.Fatal("Test GetBookWithName Error.")
	}

}

func TestGetBookWithTermName(t *testing.T) {
	ctx := context.Background()

	esConfig := config.GetEsInstance(ctx, Index)
	defer esConfig.Stop()

	books := QueryBookWithTermName(ctx, esConfig, "经济解释")
	if books == nil || len(books) == 0 {
		t.Fatal("Test GetBookWithName Error.")
	}
}

func TestCreateIndexAndMapping(t *testing.T) {
	t.Skip()
	ctx := context.Background()

	esConfig := config.GetEsInstance(ctx, Index)
	defer esConfig.Stop()

	res := esConfig.CreateIndex(ctx)
	if !res {
		t.Fatal("TestCreateIndexAndMapping Failed")
	}
	esConfig.CreateMapping(ctx, Mapping)
}

func TestSaveBooksToJsonFile(t *testing.T) {
	files, err := ioutil.ReadDir(DIR)
	if err != nil {
		t.Fatal("[*] Read Dir Failed.")
	}
    books := make([]Book, len(files))

    for i := 0; i < len(files); i++ {
        t := strings.Split(files[i].Name(), ".")
        books[i] = Book{
            ID: int64(i+1),
            Name: strings.Join(t[0:len(t)-1], "."), 
            Type: strings.ToLower(t[len(t)-1]),
            Size: files[i].Size(), 
            Created: files[i].ModTime(),
        }
    }

	file, _ := json.MarshalIndent(books, "", " ")
	_ = ioutil.WriteFile(JSONFileDir, file, 0644)
}
