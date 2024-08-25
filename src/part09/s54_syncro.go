package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func init() {
	fmt.Println("동기화 - 원자성 보장!")
}

func main54() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var cnt int64 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			//cnt += 1
			atomic.AddInt64(&cnt, 1)
			wg.Done()
		}(i)
	}

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			//cnt -= 1
			atomic.AddInt64(&cnt, -1)
			wg.Done()
		}(i)
	}

	// Add == Done 횟수 같아야
	wg.Wait() // time.Sleep() 필요없음
	finalCnt := atomic.LoadInt64(&cnt)
	fmt.Println("WaitGroup 종료!", cnt)      // 3000 으로
	fmt.Println("WaitGroup 종료!", finalCnt) // 3000
}
