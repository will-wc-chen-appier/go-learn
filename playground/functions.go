package main

import "fmt"

func sum(nums ...int) int {
	var total int
	fmt.Printf("Type: %T\n", nums)
	for _, num := range nums {
		total += num
	}

	return total
}

func checkType(values ...any) {
	for _, v := range values {
		switch x := v.(type) {
		case int:
			fmt.Println("int:", x)
		case string:
			fmt.Println("string:", x)
		case bool:
			fmt.Println("boolean:", x)
		default:
			fmt.Println("unknown type")
		}
	}
}

func variadicTest() {
	// total := sum(1, 2)
	// fmt.Println("total: ", total)
	checkType(1, "will", 3, "four", true, [2]int{4, 5})
}

func intSeq() func() int {
	x := 10
	return func() int {
		x++
		return x
	}
}

func closureTest() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInts := intSeq()
	fmt.Println(nextInts())
}

func fib(num int) int {
	if num <= 1 {
		return num
	}
	return fib(num-1) + fib(num-2)
}

func recursionTest() {
	fmt.Println(fib(3))
}
