package main

import "fmt"

func main6() {
	fmt.Println("제어문 switch")

	score := 50
	switch {
	case score >= 90:
		fmt.Println(score, "A학점")
	case score >= 75 && score < 90:
		fmt.Println(score, "B학점")
	case score >= 65 && score < 75:
		fmt.Println(score, "C학점")
	case score >= 55 && score < 65:
		fmt.Println(score, "D학점")
	default:
		fmt.Println(score, "F학점")
	}

	switch hero := "IronMan"; hero {
	case "Captain America":
		fmt.Println("스티브 로저스!")
	case "Hulk":
		fmt.Println("브루스 배너!!")
	case "AntMan":
		fmt.Println(("행크 핌!!!"))
	case "IronMan":
		fmt.Println("토니 스타크!!!!")
	default:
		fmt.Println("누군데?????")
	}
}
