package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type myElement struct {
	Name    string
	Surname string
	Id      string
}

var DATA = make(map[string]myElement)

func Add(k string, n myElement) bool {
	if k == "" {
		return false
	}
	if Lookup(k) != nil {
		return false
	}
	DATA[k] = n
	return true
}

func Delete(k string) bool {
	if Lookup(k) == nil {
		return false
	}
	delete(DATA, k)
	return true
}

func Lookup(k string) *myElement {
	v, ok := DATA[k]
	if ok {
		return &v
	}
	return nil
}

func Change(k string, n myElement) bool {
	DATA[k] = n
	return true
}

func Print() {
	for k, v := range DATA {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "PRINT":
			Print()
		case "STOP":
			return
		case "DELETE":
			if !Delete(tokens[1]) {
				fmt.Println("Delete operation failed!")
			}
		case "ADD":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !Add(tokens[1], n) {
				fmt.Println("Add operation failed!")
			}
		case "LOOKUP":
			n := Lookup(tokens[1])
			if n != nil {
				fmt.Println("%v\n", *n)
			}
		case "CHANGE":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !Change(tokens[1], n) {
				fmt.Println("Update operation failed!")
			}
		default:
			fmt.Println("usage: command <key> <field1> <field2> <field3>")
		}
	}
}
