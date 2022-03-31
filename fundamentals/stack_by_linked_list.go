package fundamentals

type stackByLinkedList[T Item] struct {
	first  *node[T]
	length int // not thread safe
}

func NewStackByLL[T Item]() *stackByLinkedList[T] {
	return &stackByLinkedList[T]{}
}

func (s *stackByLinkedList[T]) IsEmpty() bool {
	return s.first == nil
}

func (s *stackByLinkedList[T]) Size() int {
	return s.length
}

func (s *stackByLinkedList[T]) Push(t T) {
	of := s.first
	n := &node[T]{
		item: t,
		next: of,
	}
	s.first = n
	s.length += 1
}

func (s *stackByLinkedList[T]) Pop() T {
	t := s.first
	s.first = s.first.next
	s.length -= 1
	return t.item
}

func (s *stackByLinkedList[T]) Iterator() *stackLLIteractor[T] {
	return &stackLLIteractor[T]{
		s:       s,
		curr:    s.first,
		started: false,
	}
}

//////////// support for Iterator ////////////

type stackLLIteractor[T Item] struct {
	s       *stackByLinkedList[T]
	curr    *node[T]
	started bool // not thread safe
}

func (s *stackLLIteractor[T]) HasNext() bool {
	if s.started {
		return s.curr != nil
	} else {
		return s.s.first != nil
	}
}

func (s *stackLLIteractor[T]) GetNext() T {
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
