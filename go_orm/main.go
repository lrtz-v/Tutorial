package main

import (
	"fmt"
	"os"

	library "github.com/Tutorials/go_orm/library"
)

func main() {
	conn, err := library.InitDB()
	if err != nil {
		fmt.Printf("Connect...Error!\n")
		os.Exit(1)
	}

	defer conn.Close()
}
