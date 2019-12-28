package add

import "testing"

// go test -v
// go test -run TestAdd
func TestAdd(t *testing.T) {
	d := Add(1, 2)
	if d != 3 {
		t.Fatalf("[*] Except %d, Got %d.", 3, d)
	}
}

// go test -bench=.
// go test -v -bench=. -cpu=8 -benchtime="3s" -timeout="5s" -benchmem
//
// go test -bench=. -cpuprofile=cpu.out -blockprofile=block.out -memprofile=mem.out
// go tool pprof -text cpu.out
//
// go tool pprof pprof.test cpu.out
// top、top10、web
//
// go tool pprof -svg cpu.out > cpu.svg
// go tool pprof -pdf cpu.out > cpu.pdf
func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 5)
	}
}
