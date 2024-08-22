package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("채널 심화")
}

func sendOnly(ch chan<- int, cnt int) { // 발신전용
	for i := 0; i < cnt; i++ {
		ch <- i
	}
	ch <- 666
	// fmt.Println(<-ch) // 발신전용인데 수신하려고 하면 에러
}

func recvOnly(ch <-chan int) { // 수신전용
	for i := range ch {
		fmt.Println("RECV", i)
	}

	fmt.Println(<-ch)
}

func main() {
	// 채널, 함수등 매개변수에 수발신 전용 채널 지정
	// 수신 전용에 값을 넣으면 패닉 발생
	ch := make(chan int)

	go sendOnly(ch, 10) // 발신용
	go recvOnly(ch)     // 수신용

	time.Sleep(1 * time.Millisecond)
}
