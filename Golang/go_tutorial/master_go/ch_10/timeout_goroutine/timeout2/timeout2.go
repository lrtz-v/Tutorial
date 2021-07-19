package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		defer close(temp)
		time.Sleep(5 * time.Second)
		w.Wait()
	}()
	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Need a time duration!")
		return
	}

	var w sync.WaitGroup
	w.Add(1)
	t, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("Timeout period is %s\n", duration)

	if timeout(&w, duration) {
		fmt.Println("Timeout!")
	} else {
		fmt.Println("OK!")
	}

	w.Done()
	if timeout(&w, duration) {
		fmt.Println("Timeout!")
	} else {
		fmt.Println("OK!")
	}
}
