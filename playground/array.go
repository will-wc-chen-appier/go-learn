package main

import "fmt"

func arrayTest() {
	b := [...]int{100, 3: 400, 500}
	fmt.Println("idx: ", b)
	fmt.Println("idx: ", len(b))
}
