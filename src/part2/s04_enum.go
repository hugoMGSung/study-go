package main

import (
	"fmt"
)

func main4() {
	// 열거형
	// const (
	// 	Jan = 1
	// 	Feb = 2
	// 	Mar = 3
	// 	Apr = 4
	// 	May = 5
	// 	Jun = 6
	// )

	const (
		Jan = (iota + 1) * 10
		Feb
		Mar
		Apr
		May
		Jun
	)

	fmt.Println(Jan, Feb, Mar, Apr, May, Jun)
}
