package main

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Println("채널 버퍼사용/ 비동기")
}

func main38() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	ch := make(chan bool, 4) // 버퍼가 지워질때 까지
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go :", i)
			//time.Sleep(1 * time.Microsecond)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch // 수신할 때까지 대기
		fmt.Println("Main :", i)
	}
}
