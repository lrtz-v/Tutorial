package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

func increaseRandom() {
	rand.Seed(time.Now().UnixNano())
	step := 10
	maxsizemask := 1024 - 1
	i := rand.Int() & maxsizemask
	for ; step >= 0; step-- {
		fmt.Println(i)
		i = (i + 1) & maxsizemask
	}
}

func reverseCursor(v uint64) uint64 {
	CHAR_BIT := uint64(8)
	s := CHAR_BIT * uint64(unsafe.Sizeof(v))
	mask := uint64(0)
	mask = ^mask
	for {
		s >>= 1
		if s <= 0 {
			break
		}
		mask ^= (mask << s)
		v = ((v >> s) & mask) | ((v << s) & ^mask)
	}
	return v
}

func testReverseCurse(v uint64) {
	m := uint64(2048)
	v |= ^m
	v = reverseCursor(v)
	fmt.Println(v)
	v++
	fmt.Println(v)
	v = reverseCursor(v)
	fmt.Println(v)
}

func main() {
	testReverseCurse(uint64(0))
}
