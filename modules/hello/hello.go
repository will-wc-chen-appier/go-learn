package main

import (
	"fmt"
	"log"

	"examples.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{
		"Alice",
		"Bob",
		"Charlie",
		"Will",
	}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	for _, message := range messages {
		fmt.Println(message)
	}
}
