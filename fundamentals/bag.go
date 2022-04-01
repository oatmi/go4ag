package fundamentals

type Bag[T Item] struct {
	first  *node[T]
	length int
}

func NewBag[T Item]() *Bag[T] {
	return &Bag[T]{}
}

func (b *Bag[T]) Add(t T) {
	n := &node[T]{
		item: t,
		next: nil,
	}
	oldFist := b.first
	b.first = n
	b.first.next = oldFist
	b.length++
}

func (b *Bag[T]) IsEmpty() bool {
	return b.first == nil
}

func (b *Bag[T]) Size() int {
	return b.length
}

func (b *Bag[T]) Iterator() *bagIteractor[T] {
	return &bagIteractor[T]{
		s:       b,
		curr:    b.first,
		started: false,
	}
}

//////////// support for Iterator ////////////

type bagIteractor[T Item] struct {
	s       *Bag[T]
	curr    *node[T]
	started bool // not thread safe
}

func (s *bagIteractor[T]) HasNext() bool {
	if s.started {
		return s.curr != nil
	} else {
		return s.s.first != nil
	}
}

func (s *bagIteractor[T]) GetNext() T {
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
