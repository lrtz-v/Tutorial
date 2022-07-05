package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumber[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func bubbleSortByGeneric[T Number](sequence []T) {
	for i := 0; i < len(sequence)-1; i++ {
		for j := 0; j < len(sequence)-1-i; j++ {
			if sequence[j] > sequence[j+1] {
				sequence[j], sequence[j+1] = sequence[j+1], sequence[j]
			}
		}
	}
}

type Element interface {
	int64 | float64 | string
}

type Stack[V Element] struct {
	size  int
	value []V
}

func (s *Stack[V]) Push(v V) {
	s.value = append(s.value, v)
	s.size++
}

func (s *Stack[V]) Pop() V {
	e := s.value[s.size-1]
	if s.size != 0 {
		s.value = s.value[:s.size-1]
		s.size--
	}
	return e
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumber(ints),
		SumNumber(floats))

	var sequence1 = []int64{100, 23, 14, 66, 78, 12, 8}
	bubbleSortByGeneric(sequence1)
	fmt.Println(sequence1)
	var sequence2 = []float64{120.13, 2.3, 112.3, 66.5, 78.12, 1.2, 8}
	bubbleSortByGeneric(sequence2)
	fmt.Println(sequence2)

	// INT STACK
	strS := Stack[int64]{}
	strS.Push(1)
	strS.Push(2)
	strS.Push(3)
	fmt.Println(strS.Pop())
	fmt.Println(strS.Pop())
	fmt.Println(strS.Pop())

	// FLOAT STACK
	floatS := Stack[float64]{}
	floatS.Push(1.1)
	floatS.Push(2.2)
	floatS.Push(3.3)
	fmt.Println(floatS.Pop())
	fmt.Println(floatS.Pop())
	fmt.Println(floatS.Pop())
}
