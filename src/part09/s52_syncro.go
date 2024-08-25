package main

import (
	"fmt"
	"sync"
)

func init() {
	fmt.Println("동기화 - 대기그룹(중요!)")
}

func main52() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("WaitGroup", n)
			wg.Done()
		}(i)

	}

	// Add == Done 횟수 같아야
	wg.Wait() // time.Sleep() 필요없음
	fmt.Println("WaitGroup 종료!")
}
