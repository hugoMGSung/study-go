package main

import (
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("간단 웹서버2")
}

type myHandler struct {
	http.Handler
}

func main85() {
	http.Handle("/", new(myHandler))
	http.ListenAndServe(":8889", nil)
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your Request Path = " + req.URL.Path
	w.Write([]byte(str))
}
