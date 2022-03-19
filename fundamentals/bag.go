package fundamentals

type Bag[T Item] struct {
}

func NewBag[T Item]() Bag[T] {
	return Bag[T]{}
}

func (b Bag[T]) Add() {

}

func (b Bag[T]) IsEmpty() bool {
	return false
}

func (b Bag[T]) Size() int {
	return 0
}
