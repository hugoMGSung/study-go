package main

import (
	"fmt"
	"io"
	"net/http"
)

func init() {
	fmt.Println("HTTP 예제")
}

func main77() {
	req, err := http.NewRequest("GET", "https://www.yna.co.kr/", nil)
	errCheck(err)

	req.Header.Add("User-Agent", "Tester")

	client := &http.Client{}
	resp, err := client.Do(req)
	errCheck(err)

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	errCheck(err)

	fmt.Println(string(data))
}
