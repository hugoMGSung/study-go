package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		firstName string = "Hugo"
		middleName string = "MG"
		lastName string = "Sung"
		age int = 48
	)

	fmt.Println(firstName, middleName, lastName, age)

	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)
	
	i, _ := strconv.Atoi("-42")
    s := strconv.Itoa(-42)
    fmt.Println(i, s)

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum) // 5050

	// for {
	// 	fmt.Println("이건 실행 하지마셈~")
	// }
	heros := []string { "IronMan", "Captain America", "Thor", "Black Widow" }

	for index, hero := range heros {
		fmt.Println(index, hero)
	}

	name := "Hugo"
	sayHello(name)
}

func sayHello(name string) {
	fmt.Println("Hello", name)
}