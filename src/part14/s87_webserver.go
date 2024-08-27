package main

import (
	"fmt"
	"net/http"
	"time"
)

func init() {
	fmt.Println("웹서버 디렉토리 스캔-심플")
}

// 1. 첫번째 스캔
//
//	func main() {
//		port := ":8889"
//		handler := http.FileServer(http.Dir("/"))
//		http.ListenAndServe(port, handler)
//	}
//
// 2. 파일 스캔
//
//	func main() {
//		port := ":8889"
//		handler := http.StripPrefix("", http.FileServer(http.Dir("./"))) // 현재는 stripPrefix할 게 없음
//		http.Handle("/", handler)
//		http.ListenAndServe(port, nil)
//	}
//
// 3.
func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().Format("2006-01-02 15:04:05")) // 숫자를 알맞게 적어야 함
}

func main87() {
	port := ":8889"

	fileHandler := http.StripPrefix("", http.FileServer(http.Dir("./")))
	http.Handle("/", fileHandler)
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(port, nil)
}
