# Go언어

하나로 퉁치기!

## Go 문법

### 키워드
- 25개 존재
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var

### 변수와 상수

#### 변수
- 변수 선언
	var 및 type은 생략 가능

	

	```go
	var fVal float32 = 32.
	var x, y, z int
	var firstName string = ""
	firstName := "Test"
	```

	```go
	package main

	import "fmt"

	func main() {
		var (
			firstName string = "Hugo"
			middleName string = "MG"
			lastName string = "Sung"
			age int = 48
		)

		fmt.Println(firstName, middleName, lastName, age)
	}
	```

#### 상수
- 상수 선언
	```go
	const PI float32 = 3.142592
	```