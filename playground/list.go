package main

type List[T any] struct {
	head, tail *element[T]
}

func NewListFromValues[T any](vals ...T) *List[T] {
	lst := List[T]{}
	for _, val := range vals {
		lst.Push(val)
	}
	return &lst
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
