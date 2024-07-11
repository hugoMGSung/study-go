package main

import (
	"fmt"
)

func passValue(message string) {
	fmt.Printf("Total message >> [%s]\n", message)
}

func passReference(greeting *string) {
	*greeting = fmt.Sprintf("Your name is %s", *greeting)
}

func sump(nums ...int) {
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func sumr(nums ... int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	// 포인터 전달
	sayHi := "hugo"
	passReference(&sayHi)
	passValue(sayHi)

	sump(1, 2, 3, 4, 5)
	sump(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	var result = sumr(1, 2, 3, 4, 5)
	fmt.Println("sum of Return =>", result)
}