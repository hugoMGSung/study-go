package main

import (
	"fmt"
)

func main() {
	fmt.Println("for states")

	// 일반적인 사용 예
	for i := 0; i < 5; i++ {
		fmt.Println("for -", i)
	}

	// 무한루프
	var j int = 0
	for {
		j++
		fmt.Println("Infinity loop", j)
		if j == 100 {
			break
		}
	}

	// Range
	heros := []string{"Captain America", "Hulk", "AntMan", "IronMan"}

	for idx, name := range heros {
		fmt.Println("Hero", idx, name)
	}

	// 합계
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Summary", sum)

	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
	}

	sum3, i := 0, 0
	for { // 타언어 while(true)와 동일
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}

	// 후치연산은 할당이 없다. 두 변수를 사용할 때는 아래와 같이
	for i, j := 0, 0; i <= 100; i, j = 2*i+1, j+11 {
		fmt.Println("i, j", i, j)
	}

	// 구구단
	for x := 2; x < 10; x++ {
		for y := 1; y < 10; y++ {
			fmt.Print(x, "x", y, "=", x*y, "\t")
		}
		fmt.Println()
	}
}
