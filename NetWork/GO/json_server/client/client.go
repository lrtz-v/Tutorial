package main

import (
	"encoding/json"
	"fmt"
	"github.com/Tutorial/NetWork/GO/json_server/common"
	"net"
)

func main() {
	server := "127.0.0.1:1234"
	conn, err := net.Dial("tcp", server)
	common.ErrCheck(err)

	person := common.Person{
		Name: common.Name{Family: "Newmarch", Personal: "Jan"},
		Email: []common.Email{common.Email{Kind: "home", Address: "jan@newmarch.name"},
			common.Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	encoder.Encode(person)

	var newPerson common.Person
	decoder.Decode(&newPerson)
	fmt.Println(newPerson)
}
