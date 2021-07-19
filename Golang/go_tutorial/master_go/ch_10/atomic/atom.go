package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
)

type atomCounter struct {
	Val int64
}

func (c *atomCounter) Value() int64 {
	return atomic.LoadInt64(&c.Val)
}

func main() {
	minusX := flag.Int("x", 100, "Goroutines")
	minuxY := flag.Int("y", 200, "Value")
	flag.Parse()
	X := *minusX
	Y := *minuxY

	var waitGroup sync.WaitGroup
	counter := atomCounter{}

	for i := 0; i < X; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			for j := 0; j < Y; j++ {
				atomic.AddInt64(&counter.Val, 1)
			}
		}()
	}
	waitGroup.Wait()
	fmt.Println(counter.Value())

}
