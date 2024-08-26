package main

import (
	"fmt"
	"log"
)

func main62() {
	// 간단 패닉 발생
	fmt.Println("Main 시작")
	//panic("패닉 발생. 사용자 중지1")     //방법1
	log.Panic("패닉 발생. 사용자 중지2") //방법2
	fmt.Println("Main 종료")      //실행 불가
}
