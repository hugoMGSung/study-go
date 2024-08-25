package main

import (
	"fmt"
)

func init() {
	fmt.Println("채널 닫기")
}

func main40() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- 666
		}
	}()

	val1, ok1 := <-ch
	fmt.Println("ex1 : ", val1, ok1)
	val2, ok2 := <-ch
	fmt.Println("ex1 : ", val2, ok2)
	val3, ok3 := <-ch
	fmt.Println("ex1 : ", val3, ok3)

	close(ch) //채널 닫기
	val4, ok4 := <-ch
	fmt.Println("ex1 : ", val4, ok4)
}
