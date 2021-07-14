package main

import (
  "fmt"
)

func modify(arr [4]int) {
  for i, v := range arr {
    arr[i] = v+1
  }
}

func main() {
  arr := [4]int{1, 2, 3, 4}
  modify(arr)
  for _, v := range arr {
    fmt.Println(v)
  }
}
