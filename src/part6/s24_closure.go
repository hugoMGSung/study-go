package main

import (
	"fmt"
)

func init() {
	fmt.Println("클로저 기본")
}

func count_up() func() int {
	n := 0              // 지역변수(캡처됨)
	return func() int { // 익명함수
		n += 1
		return n
	}
}

func main24() {
	// 일반
	multi := func(x, y int) int {
		return x * y
	}

	res1 := multi(10, 10)
	fmt.Println(res1)

	// 클로저
	m, n := 10, 5 // 1. 변수가 캡쳐되어서
	sum := func(c int) int {
		return m + n + c // 2. 내부에서 사용된다
	}

	fmt.Println(sum(20))

	cnt := count_up()

	fmt.Println("카운트", cnt())
	fmt.Println("카운트", cnt())
	fmt.Println("카운트", cnt())
	fmt.Println("카운트", cnt())
	fmt.Println("카운트", cnt())
	fmt.Println("카운트", cnt())
	fmt.Println("카운트", cnt())
	fmt.Println("카운트", cnt())
	fmt.Println("카운트", cnt())
	fmt.Println("카운트", cnt())
}
