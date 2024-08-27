package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func init() {
	fmt.Println("웹서버 디렉토리 스캔")
}

func main86() {
	port := flag.String("p", "8889", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
