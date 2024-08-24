package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	fmt.Println("동기화 - 뮤텍스 상호배제")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	mutex := new(sync.RWMutex) // var mutex = new(sync.RWMutex)

	go func() { // 쓰기 처리
		for i := 1; i <= 10; i++ {
			mutex.Lock() // 쓰기 뮤텍스 락
			data += 1
			fmt.Println("Write", data)
			time.Sleep(200 * time.Millisecond)
			mutex.Unlock() // 쓰기 뮤텍스 잠금해제
		}
	}()

	go func() { // 읽기처리1
		for i := 1; i <= 10; i++ {
			mutex.RLock() // 읽기 뮤텍스 잠금
			fmt.Println("Read1", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock() // 읽기 뮤텍스 해제
		}
	}()

	go func() { // 읽기처리2
		for i := 1; i <= 10; i++ {
			mutex.RLock() // 읽기 뮤텍스 잠금
			fmt.Println("Read2", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock() // 읽기 뮤텍스 해제
		}
	}()

	time.Sleep(10 * time.Second)
	// 계속 10만 읽음
}
