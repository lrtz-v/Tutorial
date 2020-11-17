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

func rev(v uint64) uint64 {
	s := 8 * uint64(unsafe.Sizeof(v))
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

func dictScan(v uint64, tableSize uint64) {
	m := tableSize - 1

	fmt.Print(v)
	for {
		v |= ^m
		v = rev(v)
		v++
		v = rev(v)

		if v == 0 {
			break
		}
		fmt.Printf(" ->- (%b|%d)", v, v)
	}
	fmt.Printf("\n")
}

func main() {
	dictScan(uint64(0), uint64(8))
	// 0 ->- (100|4) ->- (10|2) ->- (110|6) ->- (1|1) ->- (101|5) ->- (11|3) ->- (111|7)

	dictScan(uint64(0), uint64(16))
	// 0 ->- (1000|8) ->- (100|4) ->- (1100|12) ->- (10|2) ->- (1010|10) ->- (110|6) ->- (1110|14) ->- (1|1) ->- (1001|9) ->- (101|5) ->- (1101|13) ->- (11|3) ->- (1011|11) ->- (111|7) ->- (1111|15)
}
