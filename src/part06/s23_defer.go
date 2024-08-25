package main

import (
	"fmt"
)

func init() {
	fmt.Println("defer 기본")
}

func call_first() {
	fmt.Println("call_first() start")
	defer run_last() //마지막에 호출
	fmt.Println("call_first() end")
}

func run_last() {
	fmt.Println("run_last() called")
}

func say_hello(name string) {
	defer func() {
		fmt.Println(name)
	}()

	func() {
		fmt.Print("Hi, ")
	}()
}

func loop() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		// defer fmt.Println(i)
	}
}

func main23() {
	call_first()

	say_hello("Hugo")

	loop()
}
