package main

import (
	"fmt"
	"strconv"
)

func init() {
	fmt.Println("메서드 기초")
}

func hi_guys() {
	fmt.Println("Hi, guys~")
}

func hello_someone(name string) {
	fmt.Println("Hello,", name)
}

func plus(x int, y int) int {
	return x + y
}

func minus(x int, y int) int {
	return x - y
}

func multiple(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	return x / y
}

func call_func(i int, f func(int, int) int) int {
	result := f(i, 10) // 콜백
	return result
}

func multi_by_value(i int) {
	i = i * 100
	// 리턴 받으면 됌!!
}

func multi_by_ref(i *int) {
	*i = *i * 200
}

func calculate(x int, y int) (int, int, int, int) {
	return x + y, x - y, x * y, x / y
}

func main21() {
	hi_guys()
	hello_someone("Hugo")

	fmt.Println(5, "+", 3, "=", plus(5, 3))
	fmt.Println(5, "-", 3, "=", minus(5, 3))
	fmt.Println(5, "x", 3, "=", multiple(5, 3))
	fmt.Println(5, "/", 3, "=", divide(5, 3))
	fmt.Println(strconv.Itoa(5), "/", strconv.Itoa(3), "=", strconv.Itoa(divide(5, 3)))

	// fmt.Println(7, "/", 0, "=", divide(7, 0)) // panic: runtime error: integer divide by zero

	fmt.Println(call_func(100, plus))

	// 값 전달
	a := 100
	multi_by_value(a)
	fmt.Println(a)

	// 참조 전달
	b := 200
	multi_by_ref(&b)
	fmt.Println(b)

	// 다중값 리턴
	p, s, m, d := calculate(10, 5)
	fmt.Println(p, s, m, d)
	fmt.Println(calculate(10, 4))
}
