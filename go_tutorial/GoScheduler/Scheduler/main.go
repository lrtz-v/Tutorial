package main

import (
	"fmt"
	"runtime"
)

func GetNumCPU() int {
	return runtime.NumCPU()
}

func main() {
	fmt.Println(GetNumCPU())
}
