package main

import (
	"fmt"
	"net/http"
)

// Just type int for the demo, can be a struct or whatever
type awesomeHandler int

// Implementing net http.Handler interface
func (h awesomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func, awesome handler")
}

func main() {
	var h awesomeHandler

	// Demonstrate that all http.Handler can be used here
	// awesomeHandler here implements ServeHTTP, so it works
	http.ListenAndServe(":8080", h)
}
