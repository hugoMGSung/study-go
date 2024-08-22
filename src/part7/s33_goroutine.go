package main

import (
	"fmt"
	"runtime"
	"time"

	"math/rand"
)

func init() {
	fmt.Println("고루틴 기초")
}

func execute4(name int) {
	r := rand.Intn(100)
	fmt.Println(name, " started : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>>", r, i)
	}
	fmt.Println(name, " func ended : ", time.Now())
}

func main33() {
	// runtime.GOMAXPROCS(runtime.NumCPU()) // 현 시스템의 CPU코어수로 설정
	runtime.GOMAXPROCS(16)
	fmt.Println("멀티코어 수", runtime.GOMAXPROCS(0))

	fmt.Println("Main Routine started", time.Now())
	for i := 0; i < 100; i++ {
		go execute4(i) // 고루틴 100개 생성 실행
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Main Routine ended", time.Now())
}
