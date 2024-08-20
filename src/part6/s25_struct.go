package main

import (
	"fmt"
)

func init() {
	fmt.Println("사용자정의 타입 기본")
}

type Car struct {
	name  string
	color string
	price int32
	tax   int32
}

type Cnt int // 기본 자료형을 사용자 정의타입으로 생성가능

type TotalCost func(int, int) int // 함수를 사용자 정의타입으로

// 일반 메서드
func Price(c Car) int32 {
	return c.price + c.tax
}

// 구조체의 메소드 바인딩
func (c Car) Price() int32 {
	return c.price + c.tax
}

func test_convert_struct(i Cnt) {
	fmt.Println("Struct", i)
}

func test_convert_type(i int) {
	fmt.Println("Integer", i)
}

func desc(price int, tax int, fnc TotalCost) {
	fmt.Println(price, ",", tax, "=", fnc(price, tax))
}

// 리시버
type shopingBasket struct{ cnt, price int }

func (b shopingBasket) purchase() int {
	return b.cnt * b.price
}

func (b *shopingBasket) rePurchaseP(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func (b shopingBasket) rePurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main25() {
	bmw := Car{name: "520d", price: 54500000, tax: 545000, color: "white"}
	benz := Car{name: "220d", price: 74500000, tax: 745000, color: "black"}

	fmt.Println(&benz == &bmw)
	fmt.Println(&bmw, &benz)
	fmt.Println(Price(bmw), Price(benz))
	fmt.Println(bmw.Price(), benz.Price())

	// 기본 자료형 변경
	a := Cnt(5)
	fmt.Println(a)
	a++
	fmt.Println(a)

	var ct Cnt = 100
	fmt.Println(ct)

	test_convert_struct(ct)
	test_convert_type(int(ct))

	// 함수 재정의
	var totalPrice TotalCost
	totalPrice = func(price int, tax int) int {
		return price + tax
	}

	desc(50000, 5000, totalPrice)

	// 리시버
	bs1 := shopingBasket{3, 5000}
	fmt.Println("bs1.purchase", bs1.purchase())
	bs1.rePurchaseP(10, 10000) //매개변수 전달(참조)
	fmt.Println("after bs1.rePurchaseP", bs1.purchase())

	bs2 := shopingBasket{5, 5000}
	fmt.Println("bs2.purchase", bs2.purchase())
	bs2.rePurchaseD(10, 10000) //매개변수 전달(복사) 원본값 수정불가
	fmt.Println("after bs2.rePurchaseD", bs2.purchase())
}
