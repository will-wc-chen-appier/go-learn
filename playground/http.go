package main

import (
	"bufio"
	"fmt"
	"net/http"
	"time"
)

func requestTest() {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i <= 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headersHandler(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func serverTest() {
	http.HandleFunc("/headers", headersHandler)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8090", nil)
}

func exampleHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Print(ctx)
	fmt.Println("server: example server started")
	defer fmt.Println("server: example server ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server: ", err)
	}
}

func contextTest() {
	http.HandleFunc("/example", exampleHandler)
	http.ListenAndServe(":8090", nil)
}
