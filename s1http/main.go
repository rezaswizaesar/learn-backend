package main

import (
	"fmt"
	"net/http"
)

func hallo() {
	fmt.Println("hello dari func")
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hallo dari handler"))
}

func main() {
	fmt.Println("hello test !!")
	hallo()

	http.HandleFunc("/app", homeHandler)

	http.ListenAndServe(":8080", nil)
}

// Terminal
// go run main.go

// scripting vs compiled language
// scripting: js, python
// compiled: go, c, c++, java

// go build .

// go run .
