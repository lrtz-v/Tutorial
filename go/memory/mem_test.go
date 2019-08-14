package main

import (
	"fmt"
	"testing"
)

func TestRound(t *testing.T) {
	var off uintptr
	var size uintptr
	off = 199
	size = 15
	if size&7 == 0 {
		off = Round(off, 8)
		fmt.Printf("8 off: %d\n", off)
	} else if size&3 == 0 {
		off = Round(off, 4)
		fmt.Printf("3 off: %d\n", off)
	} else if size&1 == 0 {
		off = Round(off, 2)
		fmt.Printf("1 off: %d\n", off)
	}

	fmt.Println("off...")
}
