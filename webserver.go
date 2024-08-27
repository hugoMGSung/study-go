package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().Format("2006-01-02 15:04:05 KST"))
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.Handle("/", http.FileServer(http.Dir("public/")))

	port := ":8080"
	log.Fatal(http.ListenAndServe(port, nil))
}
