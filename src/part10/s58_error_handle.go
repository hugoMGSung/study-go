package main

import (
	"errors"
	"fmt"
	"log"
)

func init() {
	fmt.Println("에러처리4")
}

func isZero(n int) (string, error) { // 메서드 리턴은 에러타입 중요!
	if n != 0 {
		s := fmt.Sprint("이건 0이 아님!", n)
		return s, nil
	}

	return "", errors.New("0 입력, 에러발생!")
}

func main58() {
	// Errorf 이용앟 ㄴ에러 처리
	a, err := isZero(1)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Value of a", a)

	b, err := isZero(0)

	if err != nil {
		log.Fatal(err.Error())
	}

	// Fatal 이후 프로그램 종료 후 실행 중지
	fmt.Println("Value of b", b)

	fmt.Println("Error handling 끝!")
}
