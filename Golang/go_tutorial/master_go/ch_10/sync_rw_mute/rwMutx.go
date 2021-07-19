package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

var PASSWD = secret{password: "myPassword"}

func Change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("LChange")
	time.Sleep(10 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}

func show(c *secret) string {
	c.RWM.RLock()
	fmt.Println("show")
	time.Sleep(3 * time.Second)
	defer c.RWM.RUnlock()
	return c.password
}

func showWithLock(c *secret) string {
	c.RWM.Lock()
	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)
	defer c.RWM.Unlock()
	return c.password
}

func main() {
	var showFunction = func(c *secret) string { return "" }
	if len(os.Args) != 2 {
		fmt.Println("Using sync.RWMutex!")
		showFunction = show
	} else {
		fmt.Println("Using sync.Mutex!")
		showFunction = showWithLock
	}
	var waitGroup sync.WaitGroup
	fmt.Println("Pass:", showFunction(&PASSWD))
	for i := 0; i < 15; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println("Go Pass:", showFunction(&PASSWD))
		}()
	}
	go func() {
		waitGroup.Add(1)
		defer waitGroup.Done()
		Change(&PASSWD, "123456")
	}()
	waitGroup.Wait()
	fmt.Println("Pass:", showFunction(&PASSWD))
}
