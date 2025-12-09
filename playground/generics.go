package main

import "fmt"

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// push in a new value at the tail
func (l *List[T]) Push(newVal T) {
	if l.tail == nil {
		l.head = &element[T]{val: newVal}
		l.tail = l.head
	} else {
		l.tail.next = &element[T]{val: newVal}
		l.tail = l.tail.next
	}
}

// return all the val in the list in a slice
func (l *List[T]) AllElements() []T {
	var values []T
	for e := l.head; e != nil; e = e.next {
		values = append(values, e.val)
	}
	return values
}

func genericTest() {
	// var s = []string{"foo", "bar", "zoo"}
	// _ = SlicesIndex[[]string, string](s, "zoo")
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
