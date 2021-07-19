package main

import (
	"fmt"

	"github.com/pkg/profile"
)

var VARIABLE int

func N1(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Multiply(a, b int) int {
	if a == 1 {
		return b
	}
	if a == 0 || b == 0 {
		return 0
	}
	if a < 0 {
		return -Multiply(-a, b)
	}

	return b + Multiply(a-1, b)
}

// go tool pprof -http=localhost:8080 ./cpu.pprof
func main() {
	// CPU
	defer profile.Start(profile.ProfilePath(".")).Stop()

	// MEM
	//defer profile.Start(profile.MemProfile).Stop()

	total := 0
	for i := 2; i < 200000; i++ {
		n := N1(i)
		if n {
			total++
		}
	}
	fmt.Println("Total: ", total)
	total = 0
	for i := 0; i < 5000; i++ {
		for j := 0; j < 400; j++ {
			k := Multiply(i, j)
			VARIABLE = k
			total++
		}
	}
	fmt.Println("Total: ", total)
}
