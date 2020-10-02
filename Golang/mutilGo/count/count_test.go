package count

import (
	"sync"
	"testing"
)

func TestCount(t *testing.T) {
	count(5, "🐑")
	count(5, "🐂")
}

func TestCountWithWg(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		count(5, "🐑")
		wg.Done()
	}()
	go func() {
		count(5, "🐂")
		wg.Done()
	}()
	wg.Wait()
}

func TestCountWithChannel(t *testing.T) {
	ch := make(chan string)
	go countCh(5, "🐑", ch)
	for s := range ch {
		t.Log(s)
	}
}
