package main

import "testing"

func TestRequest(t *testing.T) {
	getDate()
}

func Benchmark(b *testing.B) {
	getDate()
}

// go test -bench=.
// go test -v -bench=. -cpu=4 -benchtime="10s" -timeout="5s" -benchmem
// go test -bench=. -run=none -cpuprofile=cpu.out -blockprofile=block.out -memprofile=mem.out
// go tool pprof list  cpu.out
// go tool pprof -pdf cpu.out > cpu.pdf
// go tool pprof -http=:8080 cpu.out
