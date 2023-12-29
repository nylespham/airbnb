package main

import (
	"fmt"
	"net/http"
)

// Home is a function that returns the home page
func Home(w http.ResponseWriter, r *http.Request) {
	result, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("%d", result))
}

// About is a function that returns the about page
func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 3)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("About Page %d", sum))
}

// AddValues is a function that adds two integers
func AddValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("about", About)
	http.ListenAndServe(":8080", nil)
}
