package main

import (
	"errors"
	"fmt"
	"math"
)

func init() {
	fmt.Println("에러처리 - 심화1")
}

func Power(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, errors.New("0은 사용 불가!")
	}

	return math.Pow(f, i), nil
}

func main59() {
	// 사용자 에러 처리 생성

	if f, err := Power(7, 3); err != nil {
		fmt.Printf("Error occurred : %s\n", err.Error())
	} else {
		fmt.Printf("Power result : %.1f\n", f)
	}

	if f, err := Power(0, 3); err != nil {
		fmt.Printf("Error occurred : %s\n", err.Error())
	} else {
		fmt.Printf("Power result : %.1f\n", f)
	}
}
