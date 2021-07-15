package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func loadFromXML(filename string, key interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	decodeXml := xml.NewDecoder(f)
	err = decodeXml.Decode(key)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a filename")
		return
	}

	filename := args[1]
	var myRecord Record
	err := loadFromXML(filename, &myRecord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("XML: ", myRecord)
}
