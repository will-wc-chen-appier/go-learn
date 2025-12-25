package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func panicTest() {

	// panic("a problem")

	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}
}

func panicDemo() {
	fmt.Println("Before panic")
	defer fmt.Println("Deferred 1 - this WILL execute")
	defer fmt.Println("Deferred 2 - this WILL execute")

	panic("Something went wrong!")

	fmt.Println("After panic - this will NEVER execute")
}
