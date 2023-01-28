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

// At returns the element at the specified index.
// If the index is negative, it refers to ith element from the
// back of the deque. If index exceeds the length of the deque,
// returns ErrIndexBounds.
func (d* Deque[T]) At(index int) (item T, err error) {
	err = ErrIndexBounds
	if d == nil || d.nil == nil {
		return
	}
	if index < 0 {
		index = d.size - index
	}
	if index < 0 || index >= d.size {
		return
	}
	node := d.nil.next
	for i := 0; i < index && node != d.nil; i++ {
		node = node.next
	}
	if node == d.nil {
		// this branch shall never execute
		return
	}
	return node, nil
}
