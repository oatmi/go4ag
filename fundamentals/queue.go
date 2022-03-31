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
	q.n++
}

// Dequeue remove item from the begining of the list
func (q *Queue[T]) Dequeue() *T {
	if q.IsEmpty() {
		return nil
	}
	t := q.first
	q.first = q.first.next
	q.n--

	return &t.item
}

func (q *Queue[T]) IsEmpty() bool {
	return q.n == 0
}

func (q *Queue[T]) Size() int {
	return q.n
}

func (s *Queue[T]) Iterator() *queueIteractor[T] {
	return &queueIteractor[T]{
		s:       s,
		curr:    s.first,
		started: false,
	}
}

//////////// support for Iterator ////////////

type queueIteractor[T Item] struct {
	s       *Queue[T]
	curr    *node[T]
	started bool // not thread safe
}

func (s *queueIteractor[T]) HasNext() bool {
	if s.started {
		return s.curr != nil
	} else {
		return s.s.first != nil
	}
}

func (s *queueIteractor[T]) GetNext() T {
	if s.started {
		t := s.curr
		s.curr = s.curr.next
		return t.item
	} else {
		t := s.s.first
		s.curr = t.next
		s.started = true
		return t.item
	}
}
