// Package deque implements deque data structure for fast append and pop
// at both ends of the queue.
//
// Methods on type Deque are of the form:
//
//	d.{op}{loc}
//
// where `op` is 'Push', 'Pop' or empty;
// and `loc` is 'Front' or 'Back'
//
// The type *Deque implements Interface and internally uses a doubly linked
// list for storing items.
package deque

import "errors"

var (
	// ErrUnderflow is returned when an empty deque is queried for
	// an item.
	ErrUnderflow = errors.New("deque underflow")

	// ErrIndexBounds is returned when requested index exceeds
	// the length of the deque.
	ErrIndexBounds = errors.New("deque index out of bounds")

	// ErrInit is returned when the deque is either nil or
	// in an invalid state. This would not be needed if the
	// deque is always properly initialised.
	ErrInit = errors.New("deque is not properly initialised")
)

// Interface represents abstract data type (ADT) for deque
// data structure. It provides basic methods insert, delete
// and access elements from the deque. The type Deque implements
// Interface.
type Interface[T any] interface {
	Back() (T, error)        // item at the back of the deque
	Clear()                  // reset the deque
	Front() (T, error)       // item at the front of the deque
	Len() int                // length of the deque
	PopBack() (T, error)     // remove and return last item
	PopFront() (T, error)    // remove and return first item
	PushBack(x T)            // insert at the back of the deque
	PushFront(x T)           // insert at the front of the deque
	At(index int) (T, error) // access i-th element
}

// Must is a helper (wrapper) to cause panic in case of any error.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
