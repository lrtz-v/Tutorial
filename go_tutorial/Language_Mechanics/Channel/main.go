package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func waitForTask() {
	ch := make(chan string)

	go func() {
		_ = <-ch

		// work here.
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	ch <- "paper"
}

func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		ch <- "paper"
	}()

	_ = <-ch
}

func fanOut() {
	emps := 20
	ch := make(chan string, emps)

	for e := 0; e < emps; e++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
		}()
	}

	for emps > 0 {
		p := <-ch
		fmt.Println(p)
		emps--
	}
}

func selectDrop() {
	const cap = 5
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("employee : received :", p)
		}
	}()

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager : send ack")
		default:
			fmt.Println("manager : drop")
		}
	}

	close(ch)
}

func waitForTasks() {
	ch := make(chan string, 1)

	go func() {
		for p := range ch {
			fmt.Println("employee : working :", p)
		}
	}()

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "paper"
	}

	close(ch)
}

func withTimeout() {
	duration := 50 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch <- "paper"
	}()

	select {
	case p := <-ch:
		fmt.Println("work complete", p)

	case <-ctx.Done():
		fmt.Println("moving on")
	}
}
