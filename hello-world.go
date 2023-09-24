package main

import (
	//"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./index"))
	http.Handle("/", fs)
}
