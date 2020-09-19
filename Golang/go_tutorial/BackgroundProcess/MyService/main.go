package main

import (
	"os"
	"time"
)

const (
	FileName = "MyService.log"
)

func main() {
	for {
		appendLog()
		time.Sleep(3 * time.Second)
	}
}

func appendLog() {
	f, err := os.OpenFile(FileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	now := time.Now()
	if _, err = f.WriteString(now.String() + "\n"); err != nil {
		panic(err)
	}
}
