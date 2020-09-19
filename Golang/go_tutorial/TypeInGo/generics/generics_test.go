package generics

import "testing"

func TestPrintNumbers(t *testing.T) {
	printNumbers([]int{1, 2, 3})
	printStrings([]string{"A", "B", "C"})
}

func TestPrintAssert(t *testing.T) {
	numbers := []int{1, 2, 3}
	strings := []string{"A", "B", "C"}
	printAssert(numbers)
	printAssert(strings)
}

func TestPrintReflect(t *testing.T) {
	numbers := []int{1, 2, 3}
	strings := []string{"A", "B", "C"}
	floats := []float64{1.5, 2.9, 3.1}
	printReflect(numbers)
	printReflect(strings)
	printReflect(floats)
}

// func TestPrint(t *testing.T) {
// 	numbers := []int{1, 2, 3}
// 	strings := []string{"A", "B", "C"}
// 	floats := []float64{1.5, 2.9, 3.1}
// 	print(numbers)
// 	print(strings)
// 	print(floats)
// }
