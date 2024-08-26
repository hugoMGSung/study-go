package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func init() {
	fmt.Println("에러처리 - 심화 중요")
}

// 예외(에러) 처리 구조체
type PowError struct {
	time    time.Time //에러 발생 시간
	value   float64   //파라미터
	message string    //에러 메시지
}

func (e PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value: %g) - %s", e.time, e.value, e.message)
}

func Power3(f, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{time: time.Now(), value: f, message: "0은 사용할 수 없습니다."}
	}
	if math.IsNaN(f) {
		return 0, PowError{time: time.Now(), value: f, message: "숫자가 아닙니다."}
	}
	if math.IsNaN(i) {
		return 0, PowError{time: time.Now(), value: i, message: "숫자가 아닙니다."}
	}
	return math.Pow(f, i), nil
}

func main61() {
	v, err := Power3(10, 3) //정상 처리
	if err != nil {
		//log.Fatal(err)
		log.Fatal(err.Error())
	}

	fmt.Println("제곱값 =", v)

	//예제2
	t, err := Power3(0, 3) //예외 발생
	if err != nil {
		//log.Fatal(err)
		log.Fatal(err.Error())
		//fmt.Println(err.(PowError).value)
		//fmt.Println(err.(PowError).message)
		//fmt.Println(err.(PowError).time)
	}

	fmt.Println("제곱값 =", t)
}
