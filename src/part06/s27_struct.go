package main

import (
	"fmt"
)

func init() {
	fmt.Println("구조체 심화")
}

func NewAccount(number string, balance float64, interest float64) *Account {
	return &Account{number, balance, interest} // 구조체 인스턴스 생성 뒤
}

func CalculateD(a Account) { //값 복사 전달
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account) { //주소(참조) 전달
	a.balance = a.balance + (a.balance * a.interest)
}

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

type Executives struct {
	Employee
	specialBonus float64
}

// 이름이 같은 함수 사용 : 오버라이딩
func (e Executives) Calculate() float64 {
	return e.Employee.salary + e.Employee.bonus + e.specialBonus
}

func main27() {

	park := NewAccount("999-18-979892", 170000000, 0.01)
	fmt.Println(park)

	kim := Account{"111-11-11111", 10, 0.001}

	CalculateD(kim)
	CalculateP(&kim)
	CalculateP(park)

	fmt.Println(park)

	// 직원
	ep2 := Employee{"park", 1500000, 200000}
	// 임원
	ex := Executives{
		Employee{"lee", 4000000, 1000000},
		1000000,
	}
	fmt.Println(ep2, ex)
}
