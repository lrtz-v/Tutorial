package main

import (
	"fmt"
	"time"
)

func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	close(b)
}

func C(a chan struct{}) {
	<-a
	fmt.Println("C!")
}

func main() {
	X := make(chan struct{})
	Y := make(chan struct{})
	Z := make(chan struct{})
	go C(Z)
	go A(X, Y)
	go C(Z)
	go B(Y, Z)
	go C(Z)

	close(X)
	time.Sleep(3 * time.Second)
}
