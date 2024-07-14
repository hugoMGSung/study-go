package main

import (
	"fmt"
	"testlib"
)

// struct 정의
type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("패키지부터")

	// 패키지 사용
	song := testlib.GetMusic("Alicia Keys")
	fmt.Println(song)

	// 구조체
	psn1 := person{}

	psn1.name = "Sung"
	psn1.age = 48

	fmt.Println(psn1)

	psn2 := person{"Doug", 42}
	fmt.Println(psn2)
}
