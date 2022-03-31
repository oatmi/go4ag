package fundamentals

type Queue[T Item] struct {
	first *node[T]
	last  *node[T]
	n     int
}

func NewQueue[T Item]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue add item to the end of the list
func (q *Queue[T]) Enqueue(i T) {
	n := &node[T]{
		item: i,
		next: nil,
	}
	if q.IsEmpty() {
		q.first = n
	} else {
		q.last.next = n
	}
}

// Dequeue remove item from the begining of the list
func (q *Queue[T]) Dequeue() {

}

func (q *Queue[T]) IsEmpty() bool {
	return q.n == 0
}

func (q *Queue[T]) Size() int {
	return q.n
}
