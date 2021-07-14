package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile string
	Number string
}

func laodFromJson(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	decodeJson := json.NewDecoder(in)
	err = decodeJson.Decode(key)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a json file.")
		os.Exit(1)
	}

	filename := args[1]

	var myRecord Record
	err := laodFromJson(filename, &myRecord)
	if err != nil {
		fmt.Println(myRecord)
	} else {
		fmt.Println(err)
	}
}
