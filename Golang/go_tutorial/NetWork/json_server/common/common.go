package common

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person define person information
type Person struct {
	Name  Name
	Email []Email
}

// Name define name info
type Name struct {
	Family   string
	Personal string
}

// Email define Email Address
type Email struct {
	Kind    string
	Address string
}

// SaveJSON save object into Jsonf\File
func SaveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	ErrCheck(err)

	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	ErrCheck(err)
	outFile.Close()
}

// LoadJSON
func LoadJSON(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	ErrCheck(err)
	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(key)
	ErrCheck(err)
	inFile.Close()
}

// ErrCheck check error
func ErrCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
