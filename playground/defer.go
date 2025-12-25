package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func deferTest() {
	path := filepath.Join(os.TempDir(), "file")
	fmt.Println("File path:", path)
	file := createFile(path)
	defer closeFile(file)
	writeFile(file)
}

func createFile(path string) *os.File {
	fmt.Println("Creating")
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}

func writeFile(file *os.File) {
	fmt.Println("Writing")
	fmt.Fprintln(file, "data")
}

func closeFile(file *os.File) {
	fmt.Println("Closing")
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func deferTestBad() {
	path := filepath.Join(os.TempDir(), "file2")
	file := createFile(path)
	writeFile(file)
	closeFile(file)

	// Try to write after closing - this will fail!
	fmt.Println("Attempting to write after close...")
	_, err := fmt.Fprintln(file, "more data")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

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
