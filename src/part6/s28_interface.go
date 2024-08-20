package main

import (
	"fmt"
)

func init() {
	fmt.Println("인터페이스 기초")
}

type test interface{} // 빈 인터페이스

type Dog struct {
	name string
	year int
}

type Cat struct {
	name string
	year int
}

// Sound 메소드 구현
func (d Dog) sound() {
	fmt.Println(d.name, "왕왕!!")
}

func (d Dog) bite() {
	fmt.Println(d.name, "물어!")
}

func (d Cat) sound() {
	fmt.Println(d.name, "야옹!!")
}

func (d Cat) bite() {
	fmt.Println(d.name, "아야!!")
}

// 동물의 행동 인터페이스 선언
// 제대로 설계되었다
// 행동 인터페이스
type Behavior interface {
	sound()
	bite()
}

// 인터페이스를 파라미터로 받아 실행
func act(animal Behavior) {
	animal.sound()
	animal.bite()
}

// 익명 인터페이스
func (d Dog) run() {
	fmt.Println(d.name, "달려!!")
}

func (c Cat) run() {
	fmt.Println(c.name, "뛰어~!")
}

// 익명 인터페이스(타입을 정의하지 않음)
func new_act(animal interface{ run() }) {
	animal.run()
}

// 빈 인터페이스 활용
func printValue(s interface{}) {
	fmt.Println(s)
}

func main28() {
	var t1 test // 빈 인터페이스는 nil 리턴
	fmt.Println(t1)

	//
	dog1 := Dog{"바둑이", 5}
	var inf1 Behavior
	inf1 = dog1
	inf1.sound() // 인터페이스, Dog의 sound메서드 실행

	dog2 := Dog{"뽀삐", 10}
	inter2 := Behavior(dog2) //인터페이스 선언 및 초기화
	inter2.sound()

	inters := []Behavior{dog1, dog2} //슬라이스 타입 인터페이스 선언

	// 값 처리
	for _, dog := range inters {
		dog.sound()
	}

	// 인덱스 처리
	for idx, _ := range inters {
		inters[idx].sound()
	}

	dog := Dog{"발바리", 10}
	cat := Cat{"야옹이", 5}
	act(dog) //개 행동
	act(cat) //고양이 행동

	// 구조만 가지고 오리인지 개인지 알 수 있다. 덕타이핑
	new_act(cat)

	//빈 인터페이스 : 어떤 값이든 허용 가능(유연성 증가)
	printValue(dog)
	printValue(cat)
	printValue(15)
	printValue("Animal")
	printValue(25.5)
	printValue([]Dog{{"왕왕이", 2}, {"낑낑이", 4}})
}
