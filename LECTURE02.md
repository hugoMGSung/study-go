## Go언어 베이직

### Go의 특징

#### Go의 장점
1. 간결한 문법과 단순함(키워드 25개 끝)
2. 병렬 프로그래밍 지원, 고루틴! -> 최고 장점!!!
3. 컴파일 및 실행속도 빠름
4. 정적타입과 동적실행
5. 제네릭, 예외처리 미지원 -> 특이한 에러처리(!)
6. 간편한 협업(Git연계)
7. 컨벤션 통일(!!)

#### Go를 사용한 프로젝트
- 위키피디아에서 확인
	- https://ko.wikipedia.org/wiki/Go_(%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D_%EC%96%B8%EC%96%B4)

### Go언어 기초

#### Go의 특징
- 패키지 (fmt ...) 를 import하고 사용하지 않으면 import 구문이 자동으로 삭제
- 변수 등을 선언해놓고 사용하지 않으면 오류. 실행이 불가
- src 안에 두개의 main() 함수가 있을 수 없음
- go.mod에 패키지 이름을 적어놔야 함
- 이와 같은 이름으로 go파일에 package main 를 입력해야 함
- ++i 전치 연산은 허용하지 않음 
- Go에서 포인터를 허용하지만 복잡한 연산은 허용하지 않음


##### 변수/상수
1. 변수 var

	```go
	var i int
	var f float32 = 3.14
	var s string = "What!"
	var z = 4.56
	var g = "Hello~"
	var t = true	
	```

	- 그룹으로 선언

	```go
	var (
		firstName  string = "Hugo"
		middleName string = "MG"
		lastName   string = "Sung"
		age        int    = 48
		isOld      bool   = true
	)
	``` 

	- 짧은 선언, 함수 내에서만 사용가능, 두번 할당 불가 - 제한된 함수 범위 내에서만 사용할 경우 가독성
	```go
	localVal1 := 3
	localVal2 := false

	localVal1 := 4 // error
	```

	- go의 특이한 선언방법, i는 if문 안에서만 사용할 수 있음
	```go
	if i := 5; i < 10 {
		// ...
	}
	```
2. 상수 const
	- 한번 선언 후 값 변경 금지. 고정된 값이 필요할 경우 pi 등
	- 선언과 동시에 초기화. 나머지는 에러

##### 열거형
- 자바나 C#, python의 enum과 동일, 연속되는 값의 변수
	```go
	const (
		Jan = iota
		Feb
		Mar
		Apr
		May
		Jun
	)

	fmt.Println(Jan, Feb, Mar, Apr, May, Jun)
	// 0 1 2 3 4 5
	```

	- 값 초기화
	```go
	Jan = iota + 1
	```

	- 생략할 변수는 _ 입력
	```go
	const (
		_ = iota
		A   // 1
		_
		C   // 3
	)
	```

#### 제어문

##### if
- 반드시 bool 형으로 검사하는 상황에 따른 조건을 비교하는 문법
- 타언어는 1(true와 동일), 0(false). go에서는 사용 불가 
- 소괄호 사용안함, 중괄호는 조건 바로 뒤에

##### switch
- 앞선 if~else를 그대로 변경할 수 있음
- 짧은 선언으로 간편하게 코딩 가능
- fallthrough를 사용하면 현재 case에서 종료하지 않고 아래 case도 실행

##### for
- Go에는 반복문이 for밖에 없음
- 다양한 사용법을 숙지해야
	- 일반적인 방법
	- 무한루프
	- Range

[다음](./LECTURE03.md)