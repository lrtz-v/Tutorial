package main

import (
  "runtime"
)

type data struct {
  i, j int
}

func main() {
  var N = 40000000
  var structure []data
  for i := 0; i < N ; i++ {
    structure = append(structure, data{int(i), int(i)})
  }
  runtime.GC()
  _ = structure[0]
}
