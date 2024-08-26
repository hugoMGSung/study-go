package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("에러처리 - 심화 패닉/리커버")
}

func fOpen(filename string) {
	// defer (Panic 호출되면 실행)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("파일열기 오류!", r)
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("파일명", f.Name())
	}

	defer f.Close()
}

func main() {
	// Go/Recover 최종
	fmt.Println("Main 시작")
	fOpen("./unknown.txt")
	fmt.Println("Main 종료")
}
