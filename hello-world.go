package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	fs := http.FileServer(http.Dir("./index"))
	http.Handle("/", fs)
}
