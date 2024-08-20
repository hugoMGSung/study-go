package main

import (
	"fmt"
)

func init() {
	fmt.Println("슬라이스 기초 학습")
}

func main17() {
	// 일반 선언
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3},
		{4, 5, 6}, // 쉼표!
	} // 배열과 동일

	fmt.Println(slice1, slice2, slice3, slice4)

	if slice1 == nil {
		fmt.Println("slice1 is nil")
	}

	//slice1[0] = 10 // error
	slice4[1][1] = 9
	fmt.Println(slice4)

	fmt.Printf("%T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	// make() 사용
	var slice5 []int = make([]int, 5)
	var slice6 = make([]int, 5)
	slice7 := make([]int, 5, 10)
	slice8 := make([]int, 5)
	slice5[4] = 10 // 대입

	fmt.Printf("%T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)

	// 순회
	for i, v := range slice3 {
		fmt.Println(i, v)
	}

	// 슬라이싱(슬라이스 아님!)
	fmt.Println(slice3[:2])
}
