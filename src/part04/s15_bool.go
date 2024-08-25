package main

import (
	"fmt"
)

func init() {
	fmt.Println("불린 기초 학습")
}

func main15() {
	fmt.Println(true && true, true || false, !true)

	num1, num2 := 10, 19
	fmt.Println(num1 > num2)
	fmt.Println(num1 != num2)
}
