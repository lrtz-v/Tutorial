package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	type M [1 << 10]byte
	data := make([]*M, 1024*20)
	for i := range data {
		data[i] = new(M)
	}
	for i := range data {
		data[i] = nil
	}
}
func main() {
	test()
	for {
		var ms runtime.MemStats
		runtime.ReadMemStats(&ms)
		fmt.Printf("%s %d MB\n", time.Now().Format("15:04:05"), ms.NextGC>>20)
		time.Sleep(time.Second * 30)
	}
}

/*
go build -gcflags "-l" -o test test.go
GODEBUG="gctrace=1,gcpacertrace=1" ./test
*/
