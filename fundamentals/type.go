package fundamentals

type Item any

type node[T Item] struct {
	item T
	next *node[T]
}
