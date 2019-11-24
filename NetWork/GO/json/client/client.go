package main

import (
	"encoding/json"
	"fmt"
	"github.com/Tutorial/NetWork/GO/json/common"
	"net"
)

func main() {
	server := "127.0.0.1:1234"
	conn, err := net.Dial("tcp", server)
	ErrCheck(err)

	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	encoder.Encode(person)

	var newPerson Person
	decoder.Decode(&newPerson)
	fmt.Println(newPerson)
}
