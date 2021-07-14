package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var myTime string
	if len(os.Args) != 2 {
		fmt.Println("usage: %s string\n", filepath.Base(os.Args[0]))
	}
	myTime = os.Args[1]

	d, err := time.Parse("15:04", myTime)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Full: ", d)
	fmt.Println("Time: ", d.Hour(), d.Minute())
}
