# Deque

This package implements a double ended queue using doubly linked list.
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
memory safe. Methods are not supported on a `nil` deque.

## Examples

```go
package main

import (
	"fmt"
	"log"
	"github.com/bingxueshuang/deque"
)

func Must(s string, e error) string {
	if e != nil {
		log.Fatal(e)
		return ""
	}
	return s
}

func main() {
	q := deque.New[string]()
	q.PushBack("foo")
	q.PushBack("bar")
	q.PushFront("baz")
	
	fmt.Println("Length:", q.Len())
	fmt.Println("Front:", Must(q.Front()))
	fmt.Println("Back:", Must(q.Back()))
	
	fmt.Println("Pop:", Must(q.PopBack()))
	fmt.Println("Pop:", Must(q.PopBack()))
	fmt.Println("Pop:", Must(q.PopBack()))
}
```
