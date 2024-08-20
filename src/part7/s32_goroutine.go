package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("고루틴 기초")
}

func execute(name string) {
	fmt.Println(name, "started", time.Now())
	for i := 0; i < 1000; i++ {
		fmt.Printf("%s >>>>> %4d\n", name, i)
	}
	fmt.Println(name, "ended", time.Now())
}

func main() {
	execute("thread1")
	fmt.Println("Main Routine started", time.Now())
	go execute("thread2")
	go execute("thread3")
	time.Sleep(5 * time.Second)
	fmt.Println("Main Routine ended", time.Now())
}
