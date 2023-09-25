package main

import (
	//"fmt"
	"html/template"
	"net/http"
)

render func(w http.ResponseWriter, tmpl string, r *http.Request) {
	var tmpls = template.Must(template.ParseFiles("index/index.html", "index/2.html"))
	tmpls.Execute(w, nil)
}

func main() {



	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", h1)
	http.ListenAndServe(":8080", nil)
}
