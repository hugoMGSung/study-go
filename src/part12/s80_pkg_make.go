package main

import (
	"fmt"

	oper "main/src/part12/arithmetic"
)

func init() {
	fmt.Println("패키지 문서화")
}

func main80() {
	// 사용자 패키지 작성 및 Go 문서화
	// 기준은 GOPATH/src
	// package 패키지 이름
	//package main 을 제외하고 package 문서에 등록
	//지금까지 우리는 패키지를 사용해 왔다.
	//기본적으로 GOROOT의 패키지(pkg) 검색 -> GOPATH의 패키지(src/pkg) 검색
	//go install 명령어 실행의 경우 -> GOPATH/pkg에 등록 (문서화)
	//godoc -http=:6060(임의의포트) -> pkg 이동 -> 본인 패키지 메소드 및  주석 확인(패키지, 타입, 메소드) 주석
	nums := oper.Numbers{100, 10}
	fmt.Println("Package PLus func", nums.Plus())
	fmt.Println("Package Minus func", nums.Minus())
	fmt.Println("Package Multi func", nums.Multi())
	fmt.Println("Package Divide func", nums.Divide())

	fmt.Println("Package SquarePlus : ", nums.SquarePlus())
	fmt.Println("Package SquareMinus : ", nums.SquareMinus())
}
