package main

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("에러처리 - 심화2")
}

func Power2(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, fmt.Errorf("(%.0f) 사용불가", f)
	}

	return math.Pow(f, i), nil
}

func main60() {
	// 사용자 에러 처리 생성
	if f, err := Power2(7, 3); err != nil {
		fmt.Printf("Error occurred : %s\n", err.Error())
	} else {
		fmt.Printf("Power result : %.1f\n", f)
	}

	if f, err := Power2(0, 3); err != nil {
		fmt.Printf("Error occurred : %s\n", err.Error())
	} else {
		fmt.Printf("Power result : %.1f\n", f)
	}
}
