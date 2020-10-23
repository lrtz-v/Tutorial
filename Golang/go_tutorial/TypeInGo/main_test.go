package main

import "testing"

func TestNum(t *testing.T) {
	num1 := 1
	num2 := 1
	if num1 != num2 {
		t.Fatal("Not Equal.")
	}
}
