package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("채널 기초")
}

// 채널로 int형 수신
func working1(v chan int) {
	fmt.Println("Working1 started", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Working1 ended", time.Now())
	v <- 1 // 채널 v에 1을 전송
}

func working2(v chan int) {
	fmt.Println("Working2 started", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Working2 ended", time.Now())
	v <- 2
}

func main35() {
	fmt.Println("Main started", time.Now())

	//var c chan int
	//c = make(chan int)
	c := make(chan int) // int형 채널 선언
	go working1(c)
	go working2(c)

	<-c
	<-c
	fmt.Println("Main ended", time.Now())
	// 채널은 동기식이라 굳이 time.Sleep이 필요없다
}
