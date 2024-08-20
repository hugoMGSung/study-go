package main

import (
	"fmt"
)

func init() {
	fmt.Println("포인터 기초")
}

func callByValue(n int) {
	n = 90
}

func callByRef(n *int) {
	*n = 900
}

func main20() {
	var pointer1 *int = new(int) // basic
	var pointer2 *int            // short
	fmt.Println(pointer1, pointer2)

	var val int = 76
	pointer1 = &val
	pointer2 = &val
	fmt.Println(pointer1, pointer2)
	fmt.Println(&pointer1, &pointer2)
	fmt.Println(*pointer1, *pointer2)

	var pointer3 **int
	pointer3 = &pointer1
	fmt.Println(*pointer3, **pointer3)

	// 포인터 한번더
	pointer4 := &val
	fmt.Println(val, *pointer4)
	fmt.Println(&val, pointer4)

	*pointer4++ // 포인터의 역참조에서 1증가
	fmt.Println(val, *pointer4)

	*pointer4 = 100 // 포인터의 역참조에서 값 변경
	fmt.Println(val, *pointer4)

	val = 108 // 원본값 변경
	fmt.Println(val, *pointer4)

	// pointer4++              //컴파일 에러, 포인터 연산 허용 X
	// pointer4 = 0xc071405232 //컴파일 에러, 주소값 대입 허용 X

	var vv, rv int = 10, 10
	fmt.Println(vv, rv)

	callByValue(vv)
	callByRef(&rv)
	fmt.Println(vv, rv)
}
