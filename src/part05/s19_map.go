package main

import (
	"fmt"
)

func init() {
	fmt.Println("맵 기초")
}

func main19() {
	var map1 map[string]int = make(map[string]int) // basic init
	var map2 = make(map[string]int)                // short init

	fmt.Println(map1, map2)

	map2["Starbucks"] = 4500
	map2["Compose"] = 2000
	map2["Bacha"] = 1400000

	map3 := map[string]int{
		"Starbucks": 4500,
		"Compose":   2000,
		"Bacha":     1400000,
	}

	fmt.Println(map2)
	fmt.Println(map3)

	// 순서는 없지만
	for k, v := range map3 {
		fmt.Println(k, v)
	}

	// 추가
	map3["Mega"] = 3000
	fmt.Println(map3)

	// 삭제
	delete(map2, "Starbucks")
	fmt.Println(map2)

	// 맵조회
	price1 := map3["Bacha"]
	price2 := map3["Ediya"]
	fmt.Println(price1, price2)

	price3, res := map3["Pascucci"]
	fmt.Println(price3, res)
	price4, res := map3["Bacha"] // res로 리턴값 유무 확인
	fmt.Println(price4, res)

	// 응용
	if price, res := map3["Bacha"]; res {
		fmt.Println("가격은", price)
	} else {
		fmt.Println("가격은 몰라")
	}
}
