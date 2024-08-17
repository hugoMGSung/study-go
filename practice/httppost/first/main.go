package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 간단한 http.Post 예제
	reqBody := bytes.NewBufferString("Post plain text")
	resp, err := http.Post("https://www.naver.com/", "text/plain", reqBody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Response 체크.
	respBody, err := io.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		fmt.Println(str)
	}
}
