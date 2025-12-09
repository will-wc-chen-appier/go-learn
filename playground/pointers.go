package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func pointerTest() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("after zeroval:", i)

	zeroptr(&i)
	fmt.Println("after zeroptr:", i)

	fmt.Println("pointer: ", &i)

	fmt.Printf("type: %T\n", i)
	fmt.Printf("type: %T\n", &i)

	x := &i
	fmt.Println("pointer:", x)
	fmt.Println("value via pointer:", *x)
	fmt.Println("address of the pointer itself", &x)
}
