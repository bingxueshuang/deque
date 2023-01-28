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

// Back returns the item at the back of the deque.
// If it is called on an empty deque, it returns ErrUnderflow.
func (d *Deque[T]) Back() (item T, err error) {
	if d == nil || d.nil == nil || d.size == 0 {
		err = ErrUnderflow
		return
	}
	return d.nil.prev, nil
}

// Clear resets the deque into an empty list.
func (d *Deque[T]) Clear() {
	d.init()
}

// Front returns the item at the front of the deque.
// It returns ErrUnderflow if called on an empty deque.
func (d *Deque[T]) Front() (T, error) {
	if d == nil || d.nil == nil || d.size == 0 {
		err = ErrUnderflow
		return
	}
	return d.nil.next, nil
}

// Len returns the number of items currently stored in deque.
func (d *Deque[T]) Len() int {
	if d == nil {
		return 0
	}
	return d.size
}

// PushFront inserts item at the front of the deque.
// Returns ErrInit in case of nil deque.
func (d *Deque[T]) PushFront(item T) error {
	if d == nil {
		return ErrInit
	}
	d.insert(item, d.nil, d.nil.next)
	return nil
}

// PushBack inserts item at the back of the deque.
// Returns ErrInit in case of nil deque.
func (d *Deque[T]) PushBack(item T) error {
	if d == nil {
		return ErrInit
	}
	d.insert(item, d.nil.prev, d.nil)
	return nil
}

func (d *Deque[T]) insert(item T, prev, next *listNode[T]) {
	node := &listNode[T]{
		value: item,
		prev: prev,
		next: next,
	}
	prev.next = node
	next.prev = node
	d.size++
}

// PopFront removes and returns the item at the front of the deque.
// It returns ErrUnderflow if called on an empty deque.
func (d *Deque[T]) PopFront() (item T, err error) {
	if d == nil || d.nil == nil || d.size == 0 {
		err = ErrUnderflow
		return
	}
	d.remove(d.nil.next)
}

// PopBack removes and returns the item at the back of the deque.
// It returns ErrUnderflow if called on an empty deque.
func (d *Deque[T]) PopBack() (item T, err error) {
	if d == nil || d.nil == nil || d.size == 0 {
		err = ErrUnderflow
		return
	}
	d.remove(d.nil.prev)
}

func (d *Deque[T]) remove(node *listNode[T]) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
	d.size--
	node.prev = nil
	node.next = nil
	node = nil
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
	if index > d.size/2 {
		return d.getReverse(index)
	}
	return d.getForward(index)
}


func (d *Deque[T]) getForward(idx int) (item T, err error) {
	node := d.nil.next
	for i := 0; node != nil && i < idx && node != d.nil; i++ {
		node = node.next
	}
	if node == nil {
		err = ErrInit
		return
	}
	if node == d.nil {
		err = ErrIndexBounds
		return
	}
	return node, nil
}

func (d *Deque[T]) getReverse(idx int) (item T, err error) {
	node := d.nil.prev
	for i := 0; node != nil && i < idx && node != d.nil; i++ {
		node = node.prev
	}
	if node == nil {
		err = ErrInit
		return
	}
	if node == d.nil {
		err = ErrIndexBounds
		return
	}
	return node, nil
}
