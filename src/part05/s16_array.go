package main

import (
	"fmt"
)

func init() {
	fmt.Println("배열 기초 학습")
}

func main16() {
	var arr1 [5]int
	fmt.Println(arr1)

	var arr2 [5]int = [5]int{1, 3, 5, 7, 9}
	fmt.Printf("%7T %d\n", arr2, arr2)

	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3, len(arr3))

	arr4 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6}, //마지막 쉼표!
	}
	fmt.Printf("%T %d %v\n", arr4, len(arr4), arr4)

	// 값 입력
	arr4[1][1] = 9
	fmt.Printf("%T %d %v\n", arr4, len(arr4), arr4)

	// 값 생략
	fmt.Println()
	for i := range arr2 {
		fmt.Println("arr : ", i, arr2[i])
	}

	// 인덱스, 값
	fmt.Println()
	for i, v := range arr2 {
		fmt.Println("arr : ", i, v)
	}

	//값 복사 확인 중요
	arr5 := [5]int{1, 10, 100, 1000, 10000}
	arr6 := arr5

	fmt.Println(arr5, &arr5)
	fmt.Println(arr6, &arr6)
	fmt.Printf("%15p %v\n", &arr5, arr5) //주소 값 출력
	fmt.Printf("%p %v\n", &arr6, arr6)   //주소 값 출력

	// 정렬
	arr7 := [10]int{4, 6, 3, 9, 5, 8, 0, 7, 2, 1}
	fmt.Println(arr7)

	//sort.Ints(arr7)
	//fmt.Println(arr7)
}
