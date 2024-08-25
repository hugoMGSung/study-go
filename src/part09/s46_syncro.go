package main

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Println("동기화 - 미사용시")
}

// type count struct {
// 	num int
// }

// func (cnt *count) increment() {
// 	cnt.num += 1
// }

// func (cnt *count) result() {
// 	fmt.Println(cnt.num)
// }

func main46() {
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
