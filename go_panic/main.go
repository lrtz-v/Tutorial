package main

import "log"

func main() {
	//server()
	safelyDo()
}

func server() {
	go safelyDo()
}

func safelyDo() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	do()
}

func do() {
	panic("Do Error!")
}
