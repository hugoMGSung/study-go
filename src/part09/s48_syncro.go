package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("동기화 - 뮤텍스 미사용")
}

func main48() {
	data := 0

	go func() { // 쓰기 처리
		for i := 1; i <= 10; i++ {
			data += 1
			fmt.Println("Write", data)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() { // 읽기처리1
		for i := 1; i <= 10; i++ {
			fmt.Println("Read1", data)
			time.Sleep(1 * time.Second)

		}
	}()

	go func() { // 읽기처리2
		for i := 1; i <= 10; i++ {
			fmt.Println("Read2", data)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
	// 계속 10만 읽음
}
