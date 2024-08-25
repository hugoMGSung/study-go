package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("동기화 - 뮤텍스 사용")
}

type count struct {
	num   int
	mutex sync.Mutex
}

func (cnt *count) increment() {
	cnt.mutex.Lock() // 뮤텍스로 보호
	cnt.num += 1
	cnt.mutex.Unlock() // 공유데이터 수정 후 뮤텍스 보호 해제
}

func (cnt *count) result() {
	fmt.Println(cnt.num)
}

func main47() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cnt := count{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			cnt.increment()
			done <- true
			runtime.Gosched() // CPU 양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}

	cnt.result()
}
