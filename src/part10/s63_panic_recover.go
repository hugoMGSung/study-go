package main

import (
	"fmt"
)

func runFunc() {
	defer func() {
		s := recover()
		fmt.Println("에러메시지", s)
	}()

	panic("에러발생!")
	// fmt.Print("호출안됨")
}

func main63() {
	runFunc()

	fmt.Println("안녕!! Mie, Golang!")
}
