package deque

type listNode[T any] struct {
	value T
	prev *listNode
	next *listNode
}

type Deque[T any] struct {
	nil *listNode
	size int
}

func (d *Deque[T]) init() {
	node = &listNode[T]{}
	node.prev = node
	node.next = node
	d.nil = node
}

func New[T any]() *Deque[T] {
	deque := &Deque[T]{}
	deque.init()
	return deque
}
