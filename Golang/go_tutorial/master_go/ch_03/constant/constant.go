package main

import "fmt"

type Digit int
type Power2 int

const PI = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)

	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)
	fmt.Println(Zero)
	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println(Three)
	fmt.Println(Four)
	fmt.Println()

	const (
		p2_0 Power2 = 1 << iota
		_
		p2_2
		_
		p2_4
		_
		p2_6
	)
	fmt.Println(p2_0)
	fmt.Println(p2_2)
	fmt.Println(p2_4)
	fmt.Println(p2_6)
	fmt.Println()

	const (
		Apple, Banana, NewLine = iota + 1, iota + 2, iota + 3
		Cherimoya              = iota
		Durian
		Test
		Elderberry
		Fig
		NewLine2
	)
	fmt.Println(Apple)
	fmt.Println(Banana)
	fmt.Println(NewLine)
	fmt.Println(Cherimoya)
	fmt.Println(Durian)
	fmt.Println(Test)
	fmt.Println(Elderberry)
	fmt.Println(Fig)
	fmt.Println(NewLine2)

}
