package main

import (
	"fmt"
)

func init() {
	fmt.Println("채널 사용")
}

func rangeSum(arg int, ch chan int) {
	sum := 0

	for i := 0; i <= arg; i++ {
		sum += i
	}
	ch <- sum
}

func main36() {
	// 주로 이렇게 사용
	ch := make(chan int)

	go rangeSum(1000, ch)
	go rangeSum(100, ch)
	go rangeSum(10, ch)

	// 채널에서 값 수신완료 될때까지 대기, but!
	res1 := <-ch // 어떤게 먼저 끝나서 들어갈지는 모른다!
	res2 := <-ch
	res3 := <-ch

	fmt.Println(res1, res2, res3)
}
