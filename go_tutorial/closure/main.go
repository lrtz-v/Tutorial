package main

import (
	"fmt"
	"log"
	"time"
)

var foo6Chan = make(chan int, 10)

func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("foo1 val = %d\n", *x)
	}
}

func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo2 val = %d\n", x)
	}
}

func foo3() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fmt.Printf("foo3 val = %d\n", val)
	}
}

func foo4() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go show(val)
	}
}

// Go Routine的延迟绑定
// 闭包真正被执行的时候，for-loop结束了，但是val的生命周期在闭包内部被延长了且被赋值到最新的数值5
func foo5() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go func() {
			fmt.Printf("foo5 val = %v\n", val)
		}()
	}
}

func foo6() {
	for val := range foo6Chan {
		go func() {
			log.Printf("foo6 val = %d\n", val)
		}()
	}
}

func foo8() {
    for i := 1; i < 10; i++ {
        curTime := time.Now().UnixNano()
        go func(t1 int64, i int) {
            t2 := time.Now().UnixNano()
            fmt.Printf("foo8 ts = %d us i: %d \n", t2-t1, i)
        }(curTime, i)
    }
}

// 闭包延迟
// for-loop内部仅仅是声明了一个闭包，foo7()返回的也仅仅是一段闭包的函数定义，只有在外部执行了时才真正执行了闭包，此时才闭包内部的变量才会进行赋值的操作
func foo7(x int) []func() {
	var fs []func()
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fs = append(fs, func() {
			fmt.Printf("foo7 val = %d\n", x+val)
		})
	}
	return fs
}

func show(v interface{}) {
	fmt.Printf("foo4 val = %v\n", v)
}
