package main

import "fmt"

func genericTest() {
	// var s = []string{"foo", "bar", "zoo"}
	// _ = SlicesIndex[[]string, string](s, "zoo")
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}

// SlicesIndex takes a slice of any comparable type and an element of that type
// returns the index of the first occurrence of v in s, or -1 if not present.
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}
