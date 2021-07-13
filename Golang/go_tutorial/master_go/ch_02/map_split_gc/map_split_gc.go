package main

import (
  "runtime"
)

func main() {
  N := 40000000
  split := make([]map[int]int, 200)
  
  for i := range split {
    split[i] = make(map[int]int)
  }

  for i := 0; i < N; i++ {
    v := int(i)
    split[i%200][v] = v
  }

  runtime.GC()
  _ = split[0][0]
}
