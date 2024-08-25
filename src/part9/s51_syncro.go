package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	fmt.Println("동기화 - 심화")
}

func onceTest() {
	fmt.Println("한번만 실행될거임")
}

func main51() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once) // Once 객체 생성

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println("고루틴", n)
			once.Do(onceTest)
		}(i)
	}

	time.Sleep(1 * time.Second)
}
