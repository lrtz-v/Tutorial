package main

// #cgo CFLAGS: -I${SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include "callC.h"
import "C"

import (
  "fmt"
  "unsafe"
)

func main() {
  fmt.Println("Going to call a C function")
  C.cHello()
  fmt.Println("Going to call a C function")
  message := C.CString("This is ...")
  defer C.free(unsafe.Pointer(message))
  C.printMessage(message)
  fmt.Println("All perfectly done!")
}

