package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Address struct {
	City, Country string
}

type Employee struct {
	XMLName   xml.Name `xml:"employee"`
	Id        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Initials  string   `xml:"name>initials"`
	Height    float32  `xml:"height,omitempty"`
	Address
	Comment string `xml:",comment"`
}

func main() {

	r := &Employee{Id: 7, FirstName: "Mihalis", LastName: "Tsoukalos", Initials: "MIT"}
	r.Comment = "xxxxxx"
	r.Address = Address{"111", "222"}

	output, err := xml.MarshalIndent(r, "   ", "   ")
	if err != nil {
		fmt.Println("Error: err")
		return
	}
	output = []byte(xml.Header + string(output))
	os.Stdout.Write(output)
	os.Stdout.Write([]byte("\n"))
}
