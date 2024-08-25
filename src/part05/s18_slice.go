package main

import (
	"fmt"
	"sort"
)

func init() {
	fmt.Println("슬라이스 예제")
}

func main18() {
	slice1 := []int{1, 3, 5, 7}
	slice2 := []int{2, 4, 6, 8, 10}
	slice3 := []int{21, 22, 23, 24}
	slice1 = append(slice1, 11, 12)
	slice2 = append(slice1, slice2...)      // 슬라이스를 삽입 할 경우
	slice3 = append(slice2, slice3[0:3]...) // 추출 후 병합
	fmt.Println(slice1, slice2, slice3)

	slice5 := make([]int, 0, 5) // 다 찼을때 5씩 자동증가
	for i := 1; i <= 10; i++ {
		slice5 = append(slice5, i)
		fmt.Printf("len = %d, cap = %d, value = %v\n", len(slice5), cap(slice5), slice5)
	}
	fmt.Println(slice5)

	// 정렬
	slice6 := []int{4, 6, 3, 9, 5, 8, 0, 7, 2, 1}
	fmt.Println(slice6)
	fmt.Println(sort.IntsAreSorted(slice6))
	sort.Ints(slice6)
	fmt.Println(slice6)

	// 문자열 정렬
	slice7 := []string{"ball", "day", "fly", "apple", "cat", "egg"}
	fmt.Println(slice7)
	fmt.Println(sort.StringsAreSorted(slice7))
	sort.Strings(slice7)
	fmt.Println(slice7)

	// 복사
	slice10 := []int{1, 3, 5, 7, 9, 11}
	slice11 := make([]int, 5)
	slice12 := []int{}

	copy(slice11, slice10)
	fmt.Println(slice11)
	slice11[2] = 4
	fmt.Println(slice10, slice11) // 깊은 복사(두 객체는 독립적 메모리)
	fmt.Println(slice12)

	slice13 := slice10[0:5]
	fmt.Println(slice13)
	slice13[4] = 10
	fmt.Println(slice10, slice13) // 얕은 복사(둘 모두 바뀜)

	// 슬라이스 용량 확인
	slice14 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice15 := slice14[0:5:7] // 마지막 7은 용량(capability)
	fmt.Printf("%T, %v, %d, %d\n", slice14, slice14, len(slice14), cap(slice14))
	fmt.Printf("%T, %v, %d, %d\n", slice15, slice15, len(slice15), cap(slice15))
	slice14 = append(slice14, 11)
	slice15 = append(slice15, []int{7, 8}...)
	fmt.Printf("%T, %v, %d, %d\n", slice14, slice14, len(slice14), cap(slice14))
	fmt.Printf("%T, %v, %d, %d\n", slice15, slice15, len(slice15), cap(slice15))
}
