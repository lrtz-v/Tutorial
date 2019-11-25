package common

import (
	"fmt"
	"os"
)

const (
	DIR = "dir"
	CD  = "cd"
	PWD = "pwd"
)

// ErrCheck check error
func ErrCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
