package common

import (
	"encoding/gob"
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

// SaveGob save object into file
func SaveGob(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	ErrCheck(err)

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	ErrCheck(err)
	outFile.Close()
}

// LoadGob
func LoadGob(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	ErrCheck(err)
	decoder := gob.NewDecoder(inFile)
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
