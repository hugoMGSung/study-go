package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	fmt.Println("에러처리1")
}

func main55() {
	f, err := os.Open("unnamedfile") // 파일열기 err 발생
	if err != nil {
		log.Fatal(err.Error()) // 방법1
		//log.Fatal(err) // 방법2
	}

	fmt.Println("Filename >", f.Name())
}
