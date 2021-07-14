package main

import "fmt"

func main() {
 s1 := make([]int, 5)
 reSlice := s1[1:3]
 fmt.Println(s1)
 fmt.Println(reSlice)
 fmt.Println("s1 len: ", len(s1),  " cap: ", cap(s1))
 fmt.Println("reSlice len: ", len(reSlice), " cap: ", cap(reSlice))

 reSlice[0] = -100
 reSlice[1] = 123456
 fmt.Println(s1)
 fmt.Println(reSlice)
 fmt.Println("s1 len: ", len(s1),  " cap: ", cap(s1))
 fmt.Println("reSlice len: ", len(reSlice), " cap: ", cap(reSlice))

 reSlice = append(reSlice, 111)
 fmt.Println(s1)
 fmt.Println(reSlice)
 fmt.Println("s1 len: ", len(s1),  " cap: ", cap(s1))
 fmt.Println("reSlice len: ", len(reSlice), " cap: ", cap(reSlice))


 reSlice = append(reSlice, 222)
 reSlice = append(reSlice, 333)
 fmt.Println(s1)
 fmt.Println(reSlice)
 fmt.Println("s1 len: ", len(s1),  " cap: ", cap(s1))
 fmt.Println("reSlice len: ", len(reSlice), " cap: ", cap(reSlice))
}
