package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Person -
type Person struct {
	Name string
	Age  int
}

func main78() {
	// 간단한 http.PostForm 예제
	resp, err := http.PostForm("https://httpbin.org/post", url.Values{"Name": {"Hugo"}, "Age": {"50"}})
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

	fmt.Println("==============================")

	person := Person{"Hugo", 51}
	pbytes, _ := json.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	resp1, err := http.Post("https://httpbin.org/post", "application/json", buff)

	if err != nil {
		panic(err)
	}

	defer resp1.Body.Close()

	// Response 체크.
	data, err := io.ReadAll(resp1.Body)
	if err == nil {
		str := string(data)
		fmt.Println(str)
	}
}
