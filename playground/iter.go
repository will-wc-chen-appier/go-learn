package main

import (
	"fmt"
	"iter"
)

// from the iter package
// type Seq[V any] func(yield func(V) bool)
// type Seq2[K, V any] func(yield func(K, V) bool)

// All returns an iterator of a list
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		wrappedYield := func(v T) bool {
			fmt.Println("yield called with:", v)
			return yield(v)
		}

		for e := lst.head; e != nil; e = e.next {
			if !wrappedYield(e.val) {
				return
			}
		}
	}
}

func iterTest() {
	lst := NewListFromValues(10, 20, 30)
	// fmt.Println(lst.AllElements())

	// lst.All()(func(v int) bool {
	// 	fmt.Println(v)
	// 	return true // keep going
	// })

	for e := range lst.All() {
		fmt.Println(e)
	}
}
