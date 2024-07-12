package main

import (
	"fmt"
	"strings"
)

func passValue(message string) {
	fmt.Printf("Total message >> [%s]\n", message)
}

func passReference(greeting *string) {
	*greeting = fmt.Sprintf("Your name is %s", *greeting)
}

func sump(nums ...int) {
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func sumr(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func splitStrings(value string) (string, string) {
	temps := strings.Split(value, "|")
	return temps[0], temps[1]
}

func main() {
	fmt.Println("함수 학습")
	// 포인터 전달
	sayHi := "hugo"
	passReference(&sayHi)
	passValue(sayHi)

	// 가변 매개변수 함수
	sump(1, 2, 3, 4, 5)
	sump(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// 함수 리턴
	var result = sumr(1, 2, 3, 4, 5)
	fmt.Println("sum of Return =>", result)

	// 다중값 리턴 함수
	first, second := splitStrings("Hugo|Sung")
	fmt.Println(first, second)

	// 익명함수
	sum := func(nums ...float32) float32 {
		var middle float32 = 0.0
		for _, num := range nums {
			middle += num
		}
		return middle
	}

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// 일급함수
	add := func(i int, j int) int {
		return i + j
	}

	minus := func(i int, j int) int {
		return i - j
	}

	fmt.Println(calc(add, 10, 5))
	fmt.Println(calc(minus, 10, 5))
	// type 사용 함수 원형 지정
	fmt.Println(calcs(func(x int, y int) int { return x - y }, 10, 5))

	// 클로저
	firstNext := nextValue()
	fmt.Println("클로저1")
	fmt.Println(firstNext())
	fmt.Println(firstNext())
	fmt.Println(firstNext())
	fmt.Println(firstNext())

	secondNext := nextValue()
	fmt.Println("클로저2")
	fmt.Println(secondNext())
	fmt.Println(secondNext())
	fmt.Println(secondNext())
	fmt.Println(secondNext())
}

// Closure
func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

type calculator func(int, int) int

// calculator 원형 사용
func calcs(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}
