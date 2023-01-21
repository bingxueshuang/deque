package deque

type listNode[T any] struct {
	value T
	prev *listNode
	next *listNode
}

type Deque[T any] struct {
	head *listNode[T]
	tail *listNode[T]
	size int
}

func (d *Deque[T]) Back() (T, error) {
	if d.size == 0 {
		return T{}, ErrUnderflow
	}
	return d.tail.value
}

func (d *Deque[T]) Clear() {
	d.head = nil
	d.tail = nil
	d.size = 0
}

func (d *Deque[T]) Front() (T, error) {
	if d.size == 0 {
		return T{}, ErrUnderflow
	}
	return d.head.value
}

func (d *Deque[T]) Len() int {
	return d.size
}

func (d *Deque[T]) PopBack() (T, error) {
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
