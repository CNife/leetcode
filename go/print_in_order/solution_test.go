package print_in_order

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestFoo(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		foo := NewFoo()
		wg.Add(3)

		run := func(fn func()) {
			time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
			fn()
			wg.Done()
		}

		go run(foo.First)
		go run(foo.Second)
		go run(foo.Third)
		wg.Wait()

		//goland:noinspection SpellCheckingInspection
		if got := foo.buffer.String(); got != "firstsecondthird" {
			t.Fail()
		}
	}
}
