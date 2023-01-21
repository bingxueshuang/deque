package deque

import "errors"


var (
	ErrUnderflow   = errors.New("deque underflow")
	ErrIndexBounds = errors.New("deque index out of bounds")
)

type Interface[T any] interface {
	Back() (T, error)
	Clear()
	Front() (T, error)
	Len() int
	PopBack() (T, error)
	PopFront() (T, error)
	PushBack(x T)
	PushFront(x T)
	At(index int) (T, error)
}
