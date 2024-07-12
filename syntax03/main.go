package main

import (
	"fmt"
)

func main() {
	fmt.Println("컬렉션 학습")

	// 배열 선언, 정수타입 크기 5의 배열
	var arr1 [5]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5

	// fmt.Printf("%s %s %s %s %s", arr1[0], arr1[1], arr1[2], arr1[3], arr1[4])
	fmt.Printf("%d / %d / %d / %d / %d\n", arr1[0], arr1[1], arr1[2], arr1[3], arr1[4])

	// 배열 초기화
	var arr2 = [5]int{1, 2, 3, 4, 5}
	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2, arr3)

	// 다차원 배열
	var arr4 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, // 마지막에 , 없으면 에러
	}

	fmt.Println(arr4)
}
