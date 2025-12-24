package main

import (
	"cmp"
	"fmt"
	"slices"
)

func sortTest() {
	fruits := []string{"peach", "banana", "kiwi"}

	lenComp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	fruits1 := slices.Clone(fruits)
	slices.SortFunc(fruits1, lenComp)
	fmt.Println("fruits1 sorted by length:", fruits1)

	fruits2 := slices.Clone(fruits)
	slices.Sort(fruits2)
	fmt.Println("fruits2 sorted alphabetically:", fruits2)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println("people sorted by age:", people)
}
