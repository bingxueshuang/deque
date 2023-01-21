package deque

import "sync"

type listNode[T any] struct {
	value T
	prev *listNode
	next *listNode
}

type Deque[T any] struct {
	mu sync.Mutex
	head *listNode[T]
	tail *listNode[T]
	size int
}

func New[T any]() *Deque[T] {
	return &Deque{}
}

func (d *Deque[T]) Back() (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.size == 0 {
		return T{}, ErrUnderflow
	}
	return d.tail.value
}

func (d *Deque[T]) Clear() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.head = nil
	d.tail = nil
	d.size = 0
}

func (d *Deque[T]) Front() (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.size == 0 {
		return T{}, ErrUnderflow
	}
	return d.head.value
}

func (d *Deque[T]) Len() int {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.size
}

func (d *Deque[T]) PopBack() (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.size == 0 {
		return T{}, ErrUnderflow
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

func (d *Deque[T]) PopFront() (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.size == 0 {
		return T{}, ErrUnderflow
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

func (d *Deque[T]) PushFront(x T) {
	d.mu.Lock()
	defer d.mu.Unlock()
	node := &listNode{x}
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

func (d *Deque[T]) PushBack(x T) {
	d.mu.Lock()
	defer d.mu.Unlock()
	node := &listNode{x}
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

func (d *Deque[T]) At(index int) (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if index < 0 {
		index = d.size + index
	}
	if index < 0 {
		return T{}, ErrIndexBounds
	}
	if index > d.size/2 {
		return d.getReverse(index)
	}
	return d.getForward(index)
}

func (d *Deque[T]) getForward(idx int) (T, error) {
	node := d.head
	for i := 0; i < index && node != nil; i++ {
		node = node.next
	}
	if node == nil {
		return T{}, ErrIndexBounds
	}
	return node.value, nil
}

func (d *Deque[T]) getReverse(idx int) (T, error) {
	node := d.tail
	for i := d.size; i > idx && node != nil; i++ {
		node = node.prev
	}
	if node == nil {
		return T{}, ErrIndexBounds
	}
	return node.value, nil
}
