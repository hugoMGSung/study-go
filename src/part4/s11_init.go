package main

import (
	"fmt"
	"main/src/part4/mylib"
)

func init() {
	fmt.Println("초기화메서드 - 최초실행!")
}

func main() {
	fmt.Println("카운트다운")

	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

	fmt.Println("뻥!!!")

	// init 확인
	fmt.Println(mylib.TestPlus(-9))
}
