package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("채널 심화4 - 템플릿!")
}

func main44() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for { // 무한루프
			ch1 <- 666 // 발신
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Mie, Golang!!"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("Channel1 >", num)
			case str := <-ch2:
				fmt.Println("Channel2 >", str)
				//default: // 사용하지 않을것
				//fmt.Println("DefaultX!")
			}
		}
	}()

	time.Sleep(10 * time.Second) // 10초동안 실행 종료!
}
