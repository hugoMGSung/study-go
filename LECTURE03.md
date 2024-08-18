## Go언어 베이직

### Go언어 기초

#### 패키지
- 이전에 우리가 사용한 패키지, fmt, math/rand, time ...
	- C:\Program Files\Go\src 에 모두 위치
- 코드 구조화 및 재사용
- 결합도를 낮추고 응집도는 높이기 위한
- 패키지 단위 독립적이고 작은 단위로 개발 -> 모든 언어가 동일
- 패키지 이름 > 디렉토리 이름
- 같은 패키지 내 소스파일은 디렉토리 명을 패키지 명으로 사용
- 네이밍 규칙 > public 대문자, private 소문자 사용
- 사용자가 만든 패키지는 모두 GOPATH에 지정됨

##### 패키지 만들기
- 터미널에서 테스트할 것!
- part4\mylib 폴더 생성
- test.go 생성

	```go
	package mylib

	// public method
	func TestPlus(i int32) bool {
		return i > 0
	}
	```
- cust "main/src/part4/mylib" 로 변경하면 
	- cust.TestPlus(10) 로 바꿔서 사용가능

#### 초기화 메서드
- 패키지 로드시...
- 메인메서드 보다 먼저 1회 실행되는 초기화 함수
- init() 은 패키지에 넣어서 확인 해보는 것이 좋다
- init()은 여러개를 사용해도 무방하다

- 공식 레퍼런스 - https://go.dev/doc/effective_go#initialization 확인

<img src="./images/img005.png" width="730">


#### 기본문법

##### 수 타입
- 정수 
- 정수의 문자화 
	- 아스키코드 - https://namu.wiki/w/%EC%95%84%EC%8A%A4%ED%82%A4%20%EC%BD%94%EB%93%9C
	- 유니코드 - https://namu.wiki/w/%ED%98%84%EB%8C%80%20%ED%95%9C%EA%B8%80%EC%9D%98%20%EB%AA%A8%EB%93%A0%20%EA%B8%80%EC%9E%90/%EC%9C%A0%EB%8B%88%EC%BD%94%EB%93%9C

##### 문자열 타입
- 여기서 부터...

##### Boolean 타입
- 참거짓에 대한 데이터형