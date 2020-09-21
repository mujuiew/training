package main

import (
	"fmt"
	"net/http"
)

// HomePageHandle ...
func HomePageHandle(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

// UsersHandle ...
func UsersHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users Page")
}
func main() {
	http.HandleFunc("/", HomePageHandle)
	http.HandleFunc("/users", UsersHandle)
	http.ListenAndServe(":3000", nil)
}
