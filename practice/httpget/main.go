package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Request 객체 생성
	req, err := http.NewRequest("GET", "https://www.naver.com", nil)
	if err != nil {
		panic(err)
	}

	//필요시 헤더 추가 가능
	req.Header.Add("User-Agent", "Crawler")

	// Client객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 결과 출력
	bytes, _ := io.ReadAll(resp.Body)
	str := string(bytes) //바이트를 문자열로
	fmt.Println(str)
}
