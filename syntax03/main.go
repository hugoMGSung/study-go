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

	// 슬라이스 변수 선언
	var origin []int
	origin = []int{1, 2, 3, 4, 5} // 슬라이스에 리터럴값 지정
	origin[4] = 15

	fmt.Println(origin)

	// make() 함수 사용
	seconds := make([]int, 5, 15)           // 5(슬라이스 길이), 15(내부 배열 최대길이)
	fmt.Println(len(seconds), cap(seconds)) // len() 길이, cap() Capacity

	// 슬라이스에 별도의 길이와 용량을 지정하지 않으면, 0인 슬라이스 만듬
	// Nil Slice라고 함
	var slice []int

	if slice == nil {
		fmt.Println("Nil Slice")
	}
	fmt.Println(len(slice), cap(slice)) // 모두 0

	// 부분 슬라이스
	threes := []int{1, 2, 3, 4, 5}
	var fours = threes[2:4]
	fmt.Println(fours)

	var fifth = threes[3:]
	fmt.Println(fifth)

	// append
	threes = append(threes, 7, 9)
	fmt.Println(threes)

	// len=0, cap=3 인 슬라이스
	sliceA := make([]int, 0, 3)

	// 계속 한 요소씩 추가
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		// 슬라이스 길이와 용량 확인
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA)

	// copy()
	source := []int{1, 2, 3}
	target := make([]int, len(source), cap(source)*3)
	copy(target, source)
	fmt.Println(target)
	println(len(target), cap(target))

	// Map
	var intHt = map[int]string{}
	intHt[100] = "FIRST"
	intHt[200] = "SECOND"
	intHt[300] = "THIRD"
	fmt.Println(intHt[300])

	var stringHt = make(map[string]string)
	stringHt["ONE"] = "Hundred"
	stringHt["TWO"] = "Thousand"
	stringHt["THREE"] = "Million"
	fmt.Println(stringHt["TWO"])

	for key, value := range stringHt {
		fmt.Println(key, value)
	}
}
