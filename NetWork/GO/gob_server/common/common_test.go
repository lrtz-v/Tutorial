package common

import "testing"

func TestSaveJsonFile(t *testing.T) {

	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}
	SaveGob("person.gob", person)
}

func TestLoadJsonFile(t *testing.T) {
	var person Person
	LoadGob("person.gob", &person)
	if person.Name.Personal != "Jan" {
		t.Fatal("Json Load Erro")
	}
}
