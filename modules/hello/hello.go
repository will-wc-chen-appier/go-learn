package main

import (
	"fmt"

	"examples.com/greetings"
)

func main() {
	message := greetings.Hello("Hehe")
	fmt.Println(message)
}
