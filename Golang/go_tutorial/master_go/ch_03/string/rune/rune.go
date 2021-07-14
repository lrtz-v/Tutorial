package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	aSTring := []byte("ShangHai")
	for x, y := range aSTring {
		fmt.Println(x, y)
		fmt.Printf("Char: %c, str: %s\n", aSTring[x], string(aSTring[x]))
	}
	fmt.Println()

	const chinese = "上海市浦东新区"
	buf := []byte(chinese)
	fmt.Println("bytes =", len(buf))
	fmt.Println("runes =", utf8.RuneCount(buf))

	for i, w := 0, 0; i < len(chinese); i += w {
		runeValue, width := utf8.DecodeRuneInString(chinese[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		fmt.Printf("%c starts at byte position %d\n", runeValue, i)
		w = width
	}
	fmt.Println()
}
