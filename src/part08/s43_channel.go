package main

import (
	"fmt"
)

func init() {
	fmt.Println("채널 심화3")
}

func recvOnly2(cnt int) <-chan int { // 수신 전용채널
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
		tot <- 666
		tot <- 666
		close(tot)
	}()
	return tot
}

// 타언어에는 없는 메서드(함수)
func total(ch <-chan int) <-chan int {
	tot := make(chan int)
	go func() {
		a := 0
		for i := range ch {
			a += i
		}
		tot <- a
	}()
	return tot
}

func main43() {
	ch := recvOnly2(100) // 5050 + 666 + 666
	output := total(ch)

	fmt.Println("result =", <-output)
}
