# Go언어

하나로 퉁치기!

## Go 문법

### 키워드
- 25개 존재
	1. break
	2. case
	3. chan - 채널 선언. 값을 주고 받는 통로
	4. const
	5. continue
	6. defer - 함수 내에서 제일 마지막에 실행
	7. default
	8. else
	9. fallthrough - switch문에서 case를 만족해도 아래의 case를 실행하기 위해 사용
	10. for
	11. func
	12. go - goroutine 실행(?)
	13. goto - 특정 레이블로 이동
	14. if
	15. import
	16. interface - 메서드 집합체
	17. map - 해시테이블 자료구조
	18. package
	19. range - 컬렉션에서 각 요소의 인덱스와 값 반환
	20. return
	21. select
	22. struct
	23. switch
	24. type
	25. var

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

### 데이터형
- 데이터형 종류
	1. bool
	2. string
	3. int, int8, int16, int32, int64, uint, ... uint64, uintptr
	4. float32, float64, complex64, complex128
	5. byte, rune(= int32, 유니코드 코드포인트)

- String
	1. " " 기본 문자열
	2. ` ` 여러줄 문자열

- Type Conversion
	1. 기본적인 방법
	```go
	func main() {
		var i int = 100
		var u uint = uint(i)
		var f float32 = float32(i)  
		fmt.Println(f, u)
	
		str := "ABC"
		bytes := []byte(str)
		str2 := string(bytes)
		fmt.Println(bytes, str2)
	}
	```

	2. strconv 패키지 사용 - 패키지를 import하고 사용하지 않으면 오류발생!
	```go
	package main

	import (
		"fmt"
		"strconv"
	)

	func main() {
		i, _ := strconv.Atoi("-42")
		s := strconv.Itoa(-42)
		fmt.Println(i, s)
	}
	```

### 연산자
- 기본연산자는 다른 언어들과 동일함
	1. Bitwise 연산자
		```go
		z = (x & y) >> 2
		```

	2. 포인터 연산자
		```go
		var a int = 12
		var p = &a  // a의 주소 할당
		fmt.Println(*p) // p가 가리키는 주소의 실제값 12
		```

### 조건문

#### if
- 기본 내용 생략
- Optional Statement
	```go
	var i int = 1
	if v := i * 3; v < 10 {
		fmt.Println(v)
	}
	```

#### switch
- 여러값 비교
	```go
	func grade(score int) {
		switch {
		case score >= 90:
			fmt.Println("A")
		case score >= 80:
			fmt.Println("B")
		case score >= 70:
			fmt.Println("C")
		case score >= 60:
			fmt.Println("D")
		default:
			fmt.Println("No Hope")
		}
	}   
	```

- 타입비교
	```go
	switch v.(type) {
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	} 
	```

- fallthrough - 아래 case 도 실행

#### for
- 대표 반복문
	```go
	func main() {
		sum := 0
		for i := 1; i <= 100; i++ {
			sum += i
		}
		fmt.Println(sum) // 5050
	}
	```

- 무한루프
	```go
	func main() {
		for {
			fmt.Println("이건 실행 하지마셈~")
		}
	}
	```

- for range - 컬렉션에 적합
	```go
	heros := []string { "IronMan", "Captain America", "Thor", "Black Widow" }

	for index, hero := range heros {
		fmt.Println(index, hero)
	}
	```

- break, continue, goto 생략


### 함수
- Function, Method, Process 등등의 이름으로 부르는 코드블럭 단위 실행 유닛
	```go
	func main() {
		name := "Hugo"
		sayHello(name)
	}

	func sayHello(name string) {
		fmt.Println("Hello", name)
	}
	
	```

- 전달 방식
	- Pass By Value - 값을 전달하는 방식
	- Pass By Reference - &를 사용, 주소를 전달하는 방식

- 가변 매개변수(Variadic Arguments) 함수
	- 한 개 이상의 매개변수(인자)를 함수에 전달할 때

- 다중값 리턴(Multiple Return) 함수
	- 한번에 하나 이상의 값을 리턴할 때

- 익명(Anonymous) 함수
	- 함수 이름이 없는 함수

- 일급(First-class) 함수
	- 다른 함수의 파라미터로 전달하거나 다른 함수의 리턴값으로도 사용될 수 있는 함수.
	- 함수의 입력 매개변수나 리턴값으로 함수 자체가 사용될 수 있음

- type문을 사용한 함수 원형 정의
	- 함수의 원형을 표현 -> Delegate의 기능

- Closure
	- Closure는 함수 바깥에 있는 변수를 참조하는 함수값(function value). 이때의 함수는 바깥쪽 변수를 마치 함수 안으로 끌어들인 듯이 그 변수를 읽거나 쓸 수 있게 됨


### 컬렉션
- 배열은 연속적인 메모리에 동일한 데이터타입 데이터를 순차적으로 저장
	- 인덱스 0부터 시작
	- 0부터 시작하지 않는 것은 Database 외에는 거의 없음

- 배열 초기화
- 다차원 배열

- 슬라이스
	- var origin []int 와 같이 크기를 지정하지 않음
	- Go의 내장함수 make() 함수 이용
	- 부분 슬라이스, 파이썬과 동일
	- append() 다중 추가 가능, copy()

- Map
	- Key와 Value의 쌍으로 된 해시테이블
	