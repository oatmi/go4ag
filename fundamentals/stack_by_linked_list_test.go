package fundamentals

import (
	"fmt"
	"testing"
)

func TestStackByLL(t *testing.T) {
	stackByLL := NewStackByLL[string]()
	stackByLL.Push("hello")
	stackByLL.Push("yangtao")

	fmt.Println(stackByLL.Pop())
	fmt.Println(stackByLL.Pop())
}

func TestStackByLLIterator(t *testing.T) {
	stackByLL := NewStackByLL[string]()
	it := stackByLL.Iterator()

	stackByLL.Push("hello")
	stackByLL.Push("yangtao")

	for it.HasNext() {
		fmt.Println(it.GetNext())
	}

	fmt.Println(stackByLL.Pop())
	fmt.Println(stackByLL.Pop())
}
