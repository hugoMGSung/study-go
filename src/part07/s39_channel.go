package main

import (
	"fmt"
)

func init() {
	fmt.Println("채널 닫기")
}

func main39() {
	// Close : 채널을 닫는다. 닫힌 채널에 값을 전송하면 패닉(예외와 동일) 발생
	// Range : 채널안의 값을 꺼냄. 채널을 닫아야 반복문 종료!!! 아니면 무한대기!
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) //5회 채널에 값 전송 후 채널 닫기 // 안닫으면??
	}()

	for i := range ch { //채널에서 값을 꺼내온다.(채널이 Close 될 때까지)
		fmt.Println("Channel value", i)
	}
}
