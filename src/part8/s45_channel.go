package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("채널 심화5")
}

func main45() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			num := <-ch1 // 값 수신
			fmt.Println("Channel1", num)
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Mie, Golang" // 값 발신
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case ch1 <- 666: // 발신용도
			case str := <-ch2: // 수신용도
				fmt.Println("Channel2", str)
			}
		}
	}()

	time.Sleep(7 * time.Second)
}
