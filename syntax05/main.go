package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"runtime"
	"sync"
	"time"
)

// 사각형 구조체
type Rect struct {
	width, height float64 // int
}

// Circle 정의
type Circle struct {
	radius float64
}

// Shape 인터페이스
type Shape interface {
	area() float64
	perimeter() float64
}

// 면적구하기 메서드(함수)
func (r Rect) area1() float64 {
	r.width += 5
	return r.width * r.height
}

// Pointer Receiver
func (r *Rect) area2() float64 {
	r.width += 5
	return r.width * r.height
}

// Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	// 2개의 CPU 사용 / 속도체감 가능
	runtime.GOMAXPROCS(2)

	rect := Rect{10, 30}
	area := rect.area1()
	fmt.Println("사각형 전체 면적(Value Receiver)", rect.width, area)

	areaP := rect.area2()
	fmt.Println("사각형 전체 면적(Pointer Receiver)", rect.width, areaP)

	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)

	// 빈 인터페이스
	var x interface{}
	x = 1
	printIt(x)
	x = "Tom Cruise"
	printIt(x)

	// Type Assertion
	check := x.(string)
	fmt.Println(check) // Tom Cruise

	// Error 처리
	f, err := os.Open("./test.txt")
	if err != nil {
		switch err.(type) {
		default:
			println("No error")
		case error:
			log.Fatal(err.Error())
		}
	}
	fmt.Println(f.Name())

	// defer
	f2, err2 := os.Open("./test.txt")
	if err2 != nil {
		panic(err2)
	}

	// 마지막에 파일 close
	defer f2.Close()

	// 파일 읽기
	bytes := make([]byte, 16)
	f2.Read(bytes)
	fmt.Println(bytes)

	// 잘못된 파일명을 넣음
	openFile("./Invalid.txt")

	// recover에 의해 이 문장 실행됨
	println("Done after recover()")

	// 함수를 동기적으로 실행
	say("Sync")

	// 함수를 비동기적으로 실행
	go say("Async1")
	go say("Async2")
	go say("Async3")

	// 3초 대기
	time.Sleep(time.Second * 3)

	//// 익명함수 고루틴
	// WaitGroup 생성. 2개의 Go루틴을 기다림.
	var wait sync.WaitGroup
	wait.Add(2)

	// 익명함수를 사용한 goroutine
	go func() {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println("Hello")
	}()

	// 익명함수에 파라미터 전달
	go func(msg string) {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println(msg)
	}("Hi")

	wait.Wait() //Go루틴 모두 끝날 때까지 대기

	// 정수형 채널을 생성
	ch := make(chan int)

	go func() {
		ch <- 123 //채널로 123 전송
	}()

	var i int
	i = <-ch // 채널로부터 123 수신
	fmt.Println(i)

	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	// 위의 Go루틴이 끝날 때까지 대기
	<-done

	ch2 := make(chan int, 1)

	//수신자가 없더라도 보낼 수 있다.
	ch2 <- 101

	fmt.Println(<-ch2)
}

func openFile(fn string) {
	// defer 함수. panic 호출시 실행됨
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}

func printIt(v interface{}) {
	fmt.Println(v)
}

// 인터페이스의 수만큼 반복하면서 값을 리턴
func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area() //인터페이스 메서드 호출
		println(reflect.TypeOf(s).Name(), a)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, "***", i)
	}
}
