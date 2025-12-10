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

func errorTest() {
	for _, i := range []int{7, 42} {
		if r, e := errorIfFortyTwo(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("out of tea:", err)
			} else if errors.Is(err, ErrPower) {
				fmt.Println("power error:", err)
			} else {
				fmt.Println("some other error:", err)
			}
			continue
		}
		fmt.Println("tea is served")
	}
}
