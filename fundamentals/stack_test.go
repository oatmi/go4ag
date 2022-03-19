package fundamentals

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[string]()
	s.Push("hello")
	s.Push("world")
	s.Push("adasd")

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

func TestStackIterator(t *testing.T) {
	s := NewStack[string]()
	s.Push("hello")
	s.Push("world")
	s.Push("adasd")
	s.Push("adasd")

	it := s.Iterator()
	for it.HasNext() {
		fmt.Println(it.GetNext())
	}
}
