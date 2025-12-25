package main

import "fmt"

func recoverDemo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("About to panic...")
	panic("Oh no!")
	fmt.Println("This won't print")
}
