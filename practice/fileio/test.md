# Go언어

하나로 퉁치기!

## Go 문법

### 패키지
- C:\Program Files\Go\src 에 testlib 폴더와 music.go 를 생성
	- syntax04\main.go

### 구조체
- OOP 이전의 자료구조

### 함수/인터페이스
- 함수
	- Value receiver는 struct의 데이타를 복사(copy)하여 전달하며, Pointer receiver는 struct의 포인터만을 전달
- 인터페이스
	- 메서드 정의의 집합
	- 인터페이스 구현하는 곳에서는 인터페이스에 정의된 메서드를 모두 구현
	- 인터페이스 타입 - 빈 인터페이스. 어떤 타입이든 담을 수 있는 컨테이너
- Type Assertion
	- 타입확인
	```go
	// Assertion
	panic: interface conversion: interface {} is string, not int

	goroutine 1 [running]:
	main.main()
		d:/01_Programming/100_HugoBank/Mine/study-go/syntax05/main.go:72 +0x606
	```


### 에러처리
- 내장 타입으로 error 라는 interface 타입 존재. 이를 구현하여 커스텀 에러타입 생성가능
- 함수에서 결과와 에러를 함께 리턴하면 에러가 nil인지 체크할 수 있음

### 지연실행, 패닉, 리커버
- defer 키워드는 특정 문장 혹은 함수를 나중에 실행하게 연기
- panic() 함수는 현재 함수를 즉시 멈추고 현재 함수의 defer를 모두 실행하고 리턴
- recover() 함수는 panic 함수에 의한 패닉상태를 다시 정상으로 되돌림

### 고루틴
- Go 런타임이 관리하는 Lightweight 논리적 (혹은 가상적) 쓰레드
- "go" 키워드를 사용하여 함수를 호출하면, 런타임시 새로운 goroutine을 실행
- 익명함수 사용 가능

### 다중 CPU처리
- runtime.GOMAXPROCS(CPU수) 함수를 호출

### 채널
- 데이타를 주고 받는 통로. make() 함수를 통해 미리 생성되어야 함
- 채널 연산자 <- 을 통해 데이타를 주고 받음
- 상대편이 준비될 때까지 채널에서 대기함으로써 별도의 lock을 걸지 않고 데이타를 동기화
- Go 채널은 수신자와 송신자가 서로를 기다리는 속성으로 고루틴이 끝날때 까지 대기하게 할 수 있음

**채널파라미터, 채널닫기, 채널 Range, 채널 select**