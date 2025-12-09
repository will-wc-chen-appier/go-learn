package main

import "fmt"

func stringTest() {
	str := "Hello 好 keke"
	for i, r := range str {
		fmt.Println(i, r, string(r))
	}
	fmt.Println()

	runeSlice := []rune(str)
	for i, r := range runeSlice {
		fmt.Println(i, r, string(r))
	}
}
