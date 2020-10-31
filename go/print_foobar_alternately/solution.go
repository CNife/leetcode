package print_foobar_alternately

import (
	"bytes"
	"fmt"
)

type FooBar struct {
	n      int
	buffer bytes.Buffer
	fooOk  chan struct{}
	barOk  chan struct{}
}

func NewFooBar(n int) FooBar {
	return FooBar{
		n:     n,
		fooOk: make(chan struct{}),
		barOk: make(chan struct{}),
	}
}

func (fb *FooBar) Foo() {
	for i := 0; i < fb.n; i++ {
		if _, ok := <-fb.fooOk; ok {
			fb.print("foo")
			fb.barOk <- struct{}{}
		}
	}
}

func (fb *FooBar) Bar() {
	for i := 0; i < fb.n; i++ {
		fb.fooOk <- struct{}{}
		if _, ok := <-fb.barOk; ok {
			fb.print("bar")
		}
	}
}

func (fb *FooBar) print(s string) {
	_, _ = fmt.Fprint(&fb.buffer, s)
}
