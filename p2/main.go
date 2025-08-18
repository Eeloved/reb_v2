package main

import (
	"fmt"
	"net/http"
)

type CustomHandler struct {
	massage string
}

func (ch CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, ch.massage)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world. This is my first server!")
	}) // обработчик для корневой страницa

	http.Handle("/hello", CustomHandler{"Hello!"})
	http.Handle("/world", CustomHandler{"World"})

	http.Handle("/redirect", http.RedirectHandler("http://google.com", http.StatusMovedPermanently))

	http.ListenAndServe(":8080", nil)
}
