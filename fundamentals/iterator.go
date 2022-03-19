package fundamentals

type Iterator[T any] interface {
	HasNext() bool
	GetNext() T
}
