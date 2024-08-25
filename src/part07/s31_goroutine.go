package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("고루틴 기초")
}

func execute1() {
	fmt.Println("func1 execution started", time.Now())
	time.Sleep(1 * time.Second) // 1초 대기
	fmt.Println("func1 execution ended", time.Now())
}

func execute2() {
	fmt.Println("func2 execution started", time.Now())
	time.Sleep(1 * time.Second) // 1초 대기
	fmt.Println("func2 execution ended", time.Now())
}

func execute3() {
	fmt.Println("func3 execution started", time.Now())
	time.Sleep(1 * time.Second) // 1초 대기
	fmt.Println("func3 execution ended", time.Now())
}

func main31() {
	execute1()

	fmt.Println("Main Routine Started", time.Now())
	go execute2()
	go execute3()
	time.Sleep(4 * time.Second) // 주석처리하면 다른 스레드가 전부 실행하지 못함
	// fmt.Scanln() // 콘솔에서 테스트할 경우 시간 대기 용도 사용가능
	fmt.Println("Main Routine Ended", time.Now())
}
