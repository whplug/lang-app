package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloWorldPage)
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
