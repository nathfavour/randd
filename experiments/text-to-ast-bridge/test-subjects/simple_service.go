package main

import (
	"fmt"
	"net/http"
)

// Greeter is a simple service that greets users.
func Greeter(name string) string {
	// Original greeting logic
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, Greeter("World"))
	})
	http.ListenAndServe(":8080", nil)
}
