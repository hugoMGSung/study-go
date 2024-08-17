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

	psn3 := person{name: "Ashley", age: 40}
	fmt.Println(psn3)

	// 생성자 함수 호출로 생성
	dic := newDict()
	dic.data[1] = "One"
	dic.data[2] = "Two"

	fmt.Println(dic)
}

type dict struct {
	data map[int]string
}

// 생성자 함수
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}
