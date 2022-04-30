package fundamentals

type stack[T Item] struct {
	data []T
}

func NewStack[T Item]() *stack[T] {
	return &stack[T]{
		data: []T{},
	}
}

func (s *stack[T]) Push(i T) {
	s.data = append(s.data, i)
}

func (s *stack[T]) Pop() T {
	t := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return t
}

// Peak return the most recently insterted item
func (s *stack[T]) Peak() *T {
	if s.IsEmpty() {
		return nil
	}
	t := s.data[0]
	return &t
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *stack[T]) Size() int {
	return len(s.data)
}

func (s *stack[T]) Iterator() *stackIteractor[T] {
	return &stackIteractor[T]{
		s:     s,
		index: 0,
	}
}

//////////// support for Iterator ////////////
type stackIteractor[T Item] struct {
	s     *stack[T]
	index int
}

func (s *stackIteractor[T]) HasNext() bool {
	return len(s.s.data) > s.index
}

func (s *stackIteractor[T]) GetNext() T {
	t := s.s.data[s.index]
	s.index += 1
	return t
}
