package main

import (
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("간단 웹서버")
}

func main84() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello, Mie GoLang"))
	})

	http.ListenAndServe(":8890", nil)
}
