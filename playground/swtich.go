package main

import (
	"fmt"
	"time"
)

func switchTest() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("this is a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)

		}
	}

	whatAmI("hey")
}
