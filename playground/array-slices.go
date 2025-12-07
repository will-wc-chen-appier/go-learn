package main

import (
	"fmt"
	"reflect"
	"slices"
)

func arrayTest() {
	b := [...]int{100, 3: 400, 2: 500} // [100, 0, 500, 400]
	fmt.Println("idx: ", b)
	fmt.Println("idx: ", len(b))
}

func sliceTest() {
	var s []string
	fmt.Println("uninit: ", s, s == nil)

	s2 := make([]string, 3)
	fmt.Println("emp: ", s2, "len: ", len(s2), "cap: ", cap(s2))

	s3 := make([]string, len(s2))
	copy(s3, s2)
	fmt.Println("copy: ", s3)

	t := []string{"a", "b", "c"}
	t2 := []string{"a", "b", "c"}
	fmt.Println("equal: ", slices.Equal(t, t2))

	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("grid: ", grid)
	grid2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("equal: ", reflect.DeepEqual(grid, grid2))
}
