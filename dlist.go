package deque

import "sync"

// listNode is node of doubly linked list.
type listNode[T any] struct {
	value T
	prev  *listNode[T]
	next  *listNode[T]
}

// Deque is double ended queue implemented as a doubly linked list.
// Deque satisfies Interface. A Deque instance represents deque data
// structure that contains items of type specified by the type argument.
// The zero value of a Deque is ready to use empty deque.
type Deque[T any] struct {
	mu   sync.Mutex
	head *listNode[T]
	tail *listNode[T]
	size int
}

var _ Interface[int] = &Deque[int]{}

// New creates new empty Deque.
func New[T any]() *Deque[T] {
	return &Deque[T]{}
}

// Back returns the item at the back of the deque.
// If it is called on an empty deque, it returns ErrUnderflow.
func (d *Deque[T]) Back() (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	var zero T
	if d.size == 0 {
		return zero, ErrUnderflow
	}
	return d.tail.value, nil
}

// Clear resets the deque into an empty list.
func (d *Deque[T]) Clear() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.head = nil
	d.tail = nil
	d.size = 0
}

// Front returns the item at the front of the deque.
// It returns ErrUnderflow if called on an empty deque.
func (d *Deque[T]) Front() (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	var zero T
	if d.size == 0 {
		return zero, ErrUnderflow
	}
	return d.head.value, nil
}

// Len returns the number of items currently stored in deque.
func (d *Deque[T]) Len() int {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.size
}

// PopBack removes and returns the item at the back of the deque.
// It returns ErrUnderflow if called on an empty deque.
func (d *Deque[T]) PopBack() (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	var zero T
	if d.size == 0 {
		return zero, ErrUnderflow
	}
	node := d.tail
	val := node.value
	d.tail = node.prev
	if d.size == 1 {
		d.head = nil
	} else {
		d.tail.next = nil
	}
	node.prev = nil
	d.size--
	node = nil
	return val, nil
}

// PopFront removes and returns the item at the front of the deque.
// It returns ErrUnderflow if called on an empty deque.
func (d *Deque[T]) PopFront() (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	var zero T
	if d.size == 0 {
		return zero, ErrUnderflow
	}
	node := d.head
	val := node.value
	d.head = node.next
	if d.size == 1 {
		d.tail = nil
	} else {
		d.head.prev = nil
	}
	node.next = nil
	d.size--
	node = nil
	return val, nil
}

// PushFront inserts item at the front of the deque.
func (d *Deque[T]) PushFront(item T) {
	d.mu.Lock()
	defer d.mu.Unlock()
	node := &listNode[T]{value: item}
	node.next = d.head
	if d.size == 0 {
		d.tail = node
	} else {
		d.head.prev = node
	}
	d.head = node
	d.size++
	node = nil
}

// PushBack inserts item at the back of the deque.
func (d *Deque[T]) PushBack(item T) {
	d.mu.Lock()
	defer d.mu.Unlock()
	node := &listNode[T]{value: item}
	node.prev = d.tail
	if d.size == 0 {
		d.head = node
	} else {
		d.tail.next = node
	}
	d.tail = node
	d.size++
	node = nil
}

// At returns the element at the specified index.
// If the index is negative, it refers to ith element
// from the back of the deque. If index exceeds the length
// of the deque, returns ErrIndexBounds.
func (d *Deque[T]) At(index int) (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	var zero T
	if index < 0 {
		index = d.size + index
	}
	if index < 0 {
		return zero, ErrIndexBounds
	}
	if index > d.size/2 {
		return d.getReverse(index)
	}
	return d.getForward(index)
}

func (d *Deque[T]) getForward(idx int) (T, error) {
	var zero T
	node := d.head
	for i := 0; i < idx && node != nil; i++ {
		node = node.next
	}
	if node == nil {
		return zero, ErrIndexBounds
	}
	return node.value, nil
}

func (d *Deque[T]) getReverse(idx int) (T, error) {
	var zero T
	node := d.tail
	for i := d.size; i > idx && node != nil; i++ {
		node = node.prev
	}
	if node == nil {
		return zero, ErrIndexBounds
	}
	return node.value, nil
}
