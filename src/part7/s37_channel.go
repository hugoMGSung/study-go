package main

import (
	"fmt"
)

func init() {
	fmt.Println("채널 버퍼미사용")
}

func main37() {
	ch := make(chan bool)

	cnt := 6

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go :", i)
			// time.Sleep(5 * time.Millisecond)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch // 수신할 때까지 대기
		fmt.Println("Main :", i)
	}
}
