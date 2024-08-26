package main

import (
	"fmt"
	"io"
	"net/http"
)

func init() {
	fmt.Println("HTTP 예제")
}

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main85() {
	resp, err := http.Get("https://www.yna.co.kr/")
	errCheck(err)

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	errCheck(err)

	fmt.Println(string(data))
}
