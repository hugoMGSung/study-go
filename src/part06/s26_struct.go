package main

import (
	"fmt"
	"reflect"
)

func init() {
	fmt.Println("구조체 기본")
}

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

type car struct {
	company string
	color   string
}

type Vehicle struct {
	company string "제조사"
	name    string "차량명"
	color   string "색상"
}

func main26() {
	hugo := Account{number: "866-12-338815", balance: 31000000, interest: 0.04}
	ashely := Account{"594-12-116347", 200000000, 0.02}
	fmt.Println(hugo, ashely)
	fmt.Println(int(hugo.Calculate()), int(ashely.Calculate()))

	// 포인터 형식 선언
	var jung *Account = new(Account)
	jung.number = "944-02-369852"
	jung.balance = 5000000000
	jung.interest = 0.05

	hong := &Account{number: "245-902", balance: 15000000, interest: 0.04}

	lee := new(Account)
	lee.number = "245-903"
	lee.balance = 13000000
	lee.interest = 0.025

	fmt.Printf("%#v\n", hugo)
	fmt.Printf("%#v\n", hong)

	car1 := car{"Hyundai", "red"}
	car2 := new(car)
	car2.color, car2.company = "white", "Kia"
	car3 := &car{}
	car4 := &car{"Toyota", "yellow"}
	fmt.Println(car1, car2, car3, car4)

	mycar1 := Vehicle{company: "현대", name: "아이오닉", color: "white"}
	tag := reflect.TypeOf(Vehicle{})
	for i := 0; i < tag.NumField(); i++ {
		fmt.Println(tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}

	fmt.Println(mycar1)
}
