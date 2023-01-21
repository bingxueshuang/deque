package deque

import "errors"

var (
	ErrUnderflow = errors.New("deque underflow")
	ErrIndexBounds = errors.New("deque index out of bounds")
)

type Interface[T any] interface {
	func Back() (T, error)
	func Clear()
	func Front() (T, error)
	func Len() int
	func PopBack() (T, error)
	func PopFront() (T, error)
	func PushBack(x T)
	func PushFront(x T)
	func At(index int) (T, error)
}
