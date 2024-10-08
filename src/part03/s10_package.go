package main

import (
	"fmt"
	custom "main/src/part03/mylib"
	"os"
)

func main10() {
	var name string

	fmt.Print("이름을 입력하시오. ")
	fmt.Scanf("%s", &name)
	fmt.Fprintf(os.Stdout, "안녕하세요, %s님\n", name)

	// mylib/test.go 패키지 사용
	//fmt.Println("10은 양수인가?", mylib.TestPlus(10))

	// mylib의 Alias
	fmt.Println("10은 양수인가?", custom.TestPlus(10))
}
