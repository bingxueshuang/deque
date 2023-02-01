# Deque

This package implements a double ended queue using circular doubly linked list.
It provides means for efficient insertion and retrieval at both the ends. It
can be used as a stack as well as a queue.

Both LIFO (last in first out) operations and FIFO (first in first out)
operations are supported on the generic type `Deque`.

## Performance

Deque is a data structure that allows to add and remove items at both ends in
constant time. Thus, it is optimized for that purpose. Other operations like
access or modification to items present *somewhere* in the middle of the list
takes time linear the number of items in the deque.

Doubly linked list allows items to be appended dynamically (unlike array or
ring buffer implementation). Since it is coded in go, all operations are
memory safe. `nil` deque is readonly. It means that methods like `Len`, `Clear`,
`Front`, `Back`, `PopFront` and `PopBack` handle `nil` deques the same as empty
deque. But write or modify operations return `ErrInit`.

## Examples

```go
package main

import (
	"fmt"
	"log"
	"github.com/bingxueshuang/deque"
)

func main() {
	q := deque.New[string]()
	deque.Must(q.PushBack("foo"))  // { "foo" }
	deque.Must(q.PushBack("bar"))  // { "foo", "bar" }
	deque.Must(q.PushFront("baz")) // { "baz", "foo", "bar" }

	fmt.Println("Length:", q.Len()) // 3
	fmt.Println("Front:", deque.Must1(q.Front())) // baz
	fmt.Println("Back:", deque.Must1(q.Back()))   // bar

	fmt.Println("Pop:", deque.Must1(q.PopBack())) // bar
	fmt.Println("Pop:", deque.Must1(q.PopBack())) // foo
	fmt.Println("Pop:", deque.Must1(q.PopBack())) // baz
}
```
