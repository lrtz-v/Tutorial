package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "hello.h"
import "C"

// use C code in go
func main() {
	C.SayHello(C.CString("Hello, World\n"))
}
