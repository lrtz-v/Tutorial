package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	str := "床前明月光"
	fmt.Println("String length: ", len([]rune(str)))
	fmt.Println("Byte length: ", len(str))

	runes := utf16.Encode([]rune(str))
	fmt.Println(runes)
	ints := utf16.Decode(runes)
	fmt.Println(string(ints))
}
