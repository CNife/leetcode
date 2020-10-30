package print_in_order

import (
	"bytes"
	"fmt"
)

type Foo struct {
	buffer         bytes.Buffer
	firstFinished  chan struct{}
	secondFinished chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		firstFinished:  make(chan struct{}),
		secondFinished: make(chan struct{}),
	}
}

func (f *Foo) First() {
	_, _ = fmt.Fprint(&f.buffer, "first")
	f.firstFinished <- struct{}{}
}

func (f *Foo) Second() {
	<-f.firstFinished
	_, _ = fmt.Fprint(&f.buffer, "second")
	f.secondFinished <- struct{}{}
}

func (f *Foo) Third() {
	<-f.secondFinished
	_, _ = fmt.Fprint(&f.buffer, "third")
}
