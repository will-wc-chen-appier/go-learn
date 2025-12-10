package main

import (
	"errors"
	"fmt"
)

func errorIfFortyTwo(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	} else {
		return arg, nil
	}
}

var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg, nil
}

func errorTest() {
	// for _, i := range []int{7, 42} {
	// 	if r, e := errorIfFortyTwo(i); e != nil {
	// 		fmt.Println("f failed:", e)
	// 	} else {
	// 		fmt.Println("f worked:", r)
	// 	}
	// }

	// for i := range 5 {
	// 	if err := makeTea(i); err != nil {
	// 		if errors.Is(err, ErrOutOfTea) {
	// 			fmt.Println("out of tea:", err)
	// 		} else if errors.Is(err, ErrPower) {
	// 			fmt.Println("power error:", err)
	// 		} else {
	// 			fmt.Println("some other error:", err)
	// 		}
	// 		continue
	// 	}
	// 	fmt.Println("tea is served")
	// }

	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
