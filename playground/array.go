package main

import "fmt"

func arrayTest() {
	b := [...]int{100, 3: 400, 2: 500} // [100, 0, 500, 400]
	fmt.Println("idx: ", b)
	fmt.Println("idx: ", len(b))
}
