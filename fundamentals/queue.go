package fundamentals

type Queue[T Item] struct{}

func NewQueue[T Item]() Queue[T] {
	return Queue[T]{}
}

func (q Queue[T]) Enqueue() {

}

func (q Queue[T]) Dequeue() {

}

func (q Queue[T]) IsEmpty() bool {
	return false
}

func (q Queue[T]) Size() int {
	return 0
}
