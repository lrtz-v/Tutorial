package main

import (
	"fmt"
	"testing"
	"time"
)

func TestFool1(t *testing.T) {
	x := 133
	f1 := foo1(&x)
	f2 := foo2(x)
	f1() 
	f2()
	f1()
	f2()
	fmt.Printf("TestFool1 val = %d\n", x)

	x = 233
	f1()
	f2()
	f1()
	f2()
	fmt.Printf("TestFool1 val = %d\n", x)

	foo1(&x)()
	foo2(x)()
	foo1(&x)()
	foo2(x)()
	fmt.Printf("TestFool1 val = %d\n", x)
}

func TestFool7(t *testing.T) {
	x := 11
	f7s := foo7(x)
	for _, f7 := range f7s {
		f7()
	}
}


func TestFool3(t *testing.T) {
	foo3()
}


func TestFool4(t *testing.T) {
	foo4()
	foo4()
	foo4()
	foo4()
	foo4()
}

func TestFool5(t *testing.T) {
	foo5()
	foo5()
	foo5()
	foo5()
	foo5()
}

func TestFool6(t *testing.T) {
	go foo6()
	foo6Chan <- 1
	foo6Chan <- 2
	foo6Chan <- 3
	foo6Chan <- 5
	foo6Chan <- 11

	time.Sleep(time.Duration(1) * time.Nanosecond)
	foo6Chan <- 12
	time.Sleep(time.Duration(1) * time.Nanosecond)
	foo6Chan <- 13
	time.Sleep(time.Duration(1) * time.Nanosecond)
	foo6Chan <- 15

	// 微秒
	foo6Chan <- 21
	time.Sleep(time.Duration(1) * time.Microsecond)
	foo6Chan <- 22
	time.Sleep(time.Duration(1) * time.Microsecond)
	foo6Chan <- 23
	time.Sleep(time.Duration(1) * time.Microsecond)
	foo6Chan <- 25
	time.Sleep(time.Duration(10) * time.Second)
	
	// 毫秒
	foo6Chan <- 31
	time.Sleep(time.Duration(1) * time.Millisecond)
	foo6Chan <- 32
	time.Sleep(time.Duration(1) * time.Millisecond)
	foo6Chan <- 33
	time.Sleep(time.Duration(1) * time.Millisecond)
	foo6Chan <- 35
	time.Sleep(time.Duration(10) * time.Second)
	
	// 秒
	foo6Chan <- 41
	time.Sleep(time.Duration(1) * time.Second)
	foo6Chan <- 42
	time.Sleep(time.Duration(1) * time.Second)
	foo6Chan <- 43
	time.Sleep(time.Duration(1) * time.Second)
	foo6Chan <- 45
	time.Sleep(time.Duration(10) * time.Second)
	close(foo6Chan)
}

func TestFool8(t *testing.T) {
	foo8()
}