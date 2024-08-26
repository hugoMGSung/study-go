package main

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("에러처리3")
}

func main57() {
	// errors 패키지의 New 메서드 활용, Errorf 보다 많이 사용
	var err1 error = errors.New("Error occurred 1!")
	err2 := errors.New("Error occurred 2!")

	fmt.Println("error1 : ", err1)
	fmt.Println("error1 : ", err1.Error())

	fmt.Println("error2 : ", err2)
	fmt.Println("error2 : ", err2.Error())
}
