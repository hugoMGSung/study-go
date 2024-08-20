package main

import "fmt"

func main5() {
	fmt.Println("제어문 if")

	var a int = 20
	b := 20

	if a >= 10 {
		fmt.Println("10이상")
	}

	if b > 20 {
		fmt.Println("20초과")
	} else {
		fmt.Println("20포함 이하")
	}

	score := 75

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 75 && score < 90 {
		fmt.Println("B")
	} else if score >= 65 && score < 75 {
		fmt.Println("C")
	} else if score >= 55 && score < 65 {
		fmt.Println("D")
	} else {
		fmt.Println("F!!!!")
	}

}
