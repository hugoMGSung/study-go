package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	fmt.Println("고루틴 클로저")
}

func main34() {
	runtime.GOMAXPROCS(4) // CPU를 하나만 사용

	s := "Goroutine Closure : "

	for i := 0; i < 1000; i++ { //고루틴 1000개 실행
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i) //반복문 클로저는 일반적으로 즉시 실행 But 고루틴 클로저는 가장 나중에 실행
		// for i := 0; i < 1000; i++ { 가 끝난뒤 모두 실행
	}
	time.Sleep(5 * time.Second)
}
