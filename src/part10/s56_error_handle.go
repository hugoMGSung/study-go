package main

import (
	"fmt"
	"log"
)

func init() {
	fmt.Println("에러처리")
}

func notZero(n int) (string, error) { // 메서드 리턴은 에러타입 중요!
	if n != 0 {
		s := fmt.Sprint("이건 0이 아님!", n)
		return s, nil
	}

	return "", fmt.Errorf("%d를 입력, 에러발생!", n)
}

func main() {
	// Errorf 이용앟 ㄴ에러 처리
	a, err := notZero(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Value of a", a)

	b, err := notZero(0)

	if err != nil {
		log.Fatal(err)
	}

	// Fatal 이후 프로그램 종료 후 실행 중지
	fmt.Println("Value of b", b)

	fmt.Println("Error handling 끝!")
}
