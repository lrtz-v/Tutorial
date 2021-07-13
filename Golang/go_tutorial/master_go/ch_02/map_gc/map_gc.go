package main

import (
  "runtime"
)

func main() {
  N := 1000000
  myMap := make(map[int]*int)
  for i := 0; i < N; i++ {
    v := int(i)
    myMap[v] = &v
  }
  runtime.GC()
  _ = myMap[0]
}
