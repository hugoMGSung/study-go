package main

import (
	"fmt"
)

func init() {
	fmt.Println("채널 심화2")
}

// 리턴값이 수신용 채널인 함수
func sum(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
	}() // 항상 이렇게 쓴다

	return tot
}

func main42() {
	res := sum(100)
	fmt.Println("total =", <-res)
}
