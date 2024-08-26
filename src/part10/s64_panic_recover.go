package main

import (
	"fmt"
)

func runFunc2() {
	defer func() {
		if s := recover(); s != nil {
			fmt.Println("에러메시지", s)
		}
	}()

	a := [3]int{1, 2, 3}

	for i := 0; i < 5; i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i]) // index out of range
	}
}

func main64() {
	runFunc2()

	fmt.Println("안녕!! Mie, Golang!")
}
