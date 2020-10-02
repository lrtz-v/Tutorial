package count

import (
	"fmt"
	"log"
	"time"
)

func count(num int, str string) {
	for i := 0; i < num; i++ {
		fmt.Println(str)
		time.Sleep(time.Millisecond * 100)
	}
}

func countCh(num int, str string, ch chan string) {
	for i := 0; i < num; i++ {
		ch <- str
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

func countWithCh() {
	ch := make(chan string)
	go countCh(5, "🐑", ch)
	for s := range ch {
		log.Println(s)
	}
}

func countWithMutilCh() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		for {
			ch1 <- "🐑"
			time.Sleep(time.Millisecond * 100)
		}
	}()
	go func() {
		for {
			ch2 <- "🐂"
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		select {
		case msg:= <-ch1:
			fmt.Println(msg)
		case msg:= <-ch2:
			fmt.Println(msg)
		}
	}
}
