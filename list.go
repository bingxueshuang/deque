package deque

// listNode is node of circular doubly linked list
type listNode[T any] struct {
	value T
	prev *listNode
	next *listNode
}

// Deque is double ended queue data structure implemented
// using circular doubly linked list. Deque satisfies Interface.
type Deque[T any] struct {
	nil *listNode
	size int
}

// init fixes the deque by (re)initialising the deque sentinel
// with proper list node. This should be called when sentinel
// is nil.
func (d *Deque[T]) init() {
	node = &listNode[T]{}
	node.prev = node
	node.next = node
	d.nil = node
}

// New creates and initialises an empty deque ready to use.
func New[T any]() *Deque[T] {
	deque := &Deque[T]{}
	deque.init()
	return deque
}
