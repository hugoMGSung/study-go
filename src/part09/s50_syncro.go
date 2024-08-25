package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	fmt.Println("동기화 - 실행흐름 제어")
}

func main50() {
	// 동기화 상태(조건) 메소드 사용 - 실행흐름 제어
	// Wait , notify , notifyAll : Java 등
	// Wait , Signal(하나만 깨움), Broadcast(다 깨움)
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			ch <- 666
			fmt.Println("고루틴 대기", n)
			condition.Wait() // 고루틴 대기
			fmt.Println("대기종료", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		//<-ch // 또는 아래
		fmt.Println("수신", <-ch) // 값 확인
	}

	// 깨우기를 빼고 실행해볼 것
	// for i := 0; i < 5; i++ {
	// 	mutex.Lock()
	// 	fmt.Println("고루틴 깨우기(Signal)", i)
	// 	condition.Signal() //한 개 씩 깨움(모든 고루틴 생성 후)
	// 	mutex.Unlock()
	// }

	mutex.Lock()
	fmt.Println("고루틴 전부깨우기(Broadcast)")
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(2 * time.Second)
}
