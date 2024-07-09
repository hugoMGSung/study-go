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