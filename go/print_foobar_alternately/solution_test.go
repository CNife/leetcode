package print_foobar_alternately

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestFooBar(t *testing.T) {
	n := 100
	fb := NewFooBar(n)

	var wg sync.WaitGroup
	wg.Add(2)
	run := func(f func()) {
		sleepTime := rand.Int63n(100)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		f()
		wg.Done()
	}
	go run(fb.Foo)
	go run(fb.Bar)
	wg.Wait()

	got := fb.buffer.Bytes()
	want := []byte("foobar")
	if len(got) != 6*n {
		t.FailNow()
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 6; j++ {
			if got[i*6+j] != want[j] {
				t.Fatalf("error at %v", i*6+j)
			}
		}
	}
}
