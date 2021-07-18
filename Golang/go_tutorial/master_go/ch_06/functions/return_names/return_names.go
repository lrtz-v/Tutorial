package main

import (
	"fmt"
	"os"
	"strconv"
)

func nameMinMax(x, y int) (min, max int) {
	if x > y {
		max, min = x, y
	} else {
		max, min = y, x
	}
	return
}

func minMax(x, y int) (min, max int) {
	if x > y {
		max, min = x, y
	} else {
		max, min = y, x
	}
	return min, max
}

func main() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("The program needs at least 2 arguments!")
		return
	}
	a1, _ := strconv.Atoi(arguments[1])
	a2, _ := strconv.Atoi(arguments[2])
	fmt.Println(minMax(a1, a2))
	min, max := minMax(a1, a2)
	fmt.Println(min, max)
	fmt.Println(nameMinMax(a1, a2))
	min, max = nameMinMax(a1, a2)
	fmt.Println(min, max)
}
