package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/Tutorial/NetWork/GO/gob_server/common"
)

func main() {
	server := "127.0.0.1:1234"
	conn, err := net.Dial("tcp", server)
	common.ErrCheck(err)

	person := common.Person{
		Name: common.Name{Family: "Newmarch", Personal: "Jan"},
		Email: []common.Email{common.Email{Kind: "home", Address: "jan@newmarch.name"},
			common.Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	encoder.Encode(person)

	var newPerson common.Person
	decoder.Decode(&newPerson)
	fmt.Println(newPerson)
}
