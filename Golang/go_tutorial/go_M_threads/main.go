package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, i int) {
	time.Sleep(10 * time.Second)
	fmt.Println(i)
	wg.Done()
}

func main() {

	startAt := time.Now()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}

	wg.Wait()
	fmt.Println(time.Since(startAt))

}
