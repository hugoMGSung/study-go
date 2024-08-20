package main

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("메서드 응용")
}

func multi(n ...int) int {
	total := 1
	for _, value := range n {
		total *= value
	}

	return total
}

func multi2(n ...int) float64 {
	total := 1.0
	for _, value := range n {
		total *= float64(value)
	}

	return total
}

func sum(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}

	return total
}

func proc1(x, y int) (r int) {
	r = x + y
	return r
}

func proc2(x, y int) (r int) {
	r = x - y
	return r
}

// 팩토리얼
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main22() {
	x := multi(1, 3, 5, 7, 9, 11, 13, 15, 17, 19)
	fmt.Println(x)

	y := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(y)

	// 배열/슬라이스를 넘기는 방법
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 11; i <= 100; i++ {
		slice = append(slice, i)
	}

	fmt.Println(multi(slice...))

	// 함수를 슬라이스 변수에 할당
	fncs := []func(int, int) int{proc1, proc2}

	res1, res2 := fncs[0](10, 10), fncs[1](10, 5)
	fmt.Println(res1, res2)

	// 일반변수에 할당
	var fnc2 func(int, int) int = proc1
	fnc3 := proc2 // 제일 간단
	fmt.Println(fnc2(10, 10), fnc3(10, 5))

	// 맵에 할당
	mp := map[string]func(int, int) int{
		"plus":  proc1,
		"minus": proc2,
	}
	fmt.Println(mp["plus"](10, 10), mp["minus"](10, 5))

	// 재귀함수
	result := fact(10)
	fmt.Println(result)

	// 익명함수
	func() {
		fmt.Println("Anonymous function")
	}()

	text := "blah blah"
	func(msg string) {
		fmt.Println("What?!", msg)
	}(text)

	res := func(x, y float64) float64 {
		return math.Pow(x, y)
	}(2, 10)
	fmt.Println(res)
}
