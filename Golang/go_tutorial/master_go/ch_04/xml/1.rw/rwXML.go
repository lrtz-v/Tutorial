package main

import (
	"encoding/json"
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

func loadFromJson(filename string, key interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	decodeJFON := json.NewDecoder(f)
	err = decodeJFON.Decode(key)
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
	err := loadFromJson(filename, &myRecord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("JSON: ", myRecord)

	myRecord.Name = "Dimitris"

	xmlData, _ := xml.MarshalIndent(myRecord, "", "    ")
	xmlData = []byte(xml.Header + string(xmlData))
	fmt.Println("\nxmlData: ", string(xmlData))
	fmt.Println()

	data := &Record{}
	err = xml.Unmarshal(xmlData, data)
	if err != nil {
		fmt.Println("Unmarshal from XML failed", err)
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling to JSON", err)
		return
	}

	_ = json.Unmarshal([]byte(result), &myRecord)
	fmt.Println("\nJSON: ", myRecord)
}
