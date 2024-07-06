# study-go
Go언어 하나로 퉁치기!

## Go 개요

### Go란
- Go Lang이라고도 부르며 2009년 구글 개발자들이 개발시작, 2012년에 발표한 새로운 프로그래밍 언어
	- 2024년 현재, 1.22.4 버전
	- 개발자 로버트 그리즈머, 롭 파이크, 켄 톰슨
	- 켄 톰슨은 C언어를 개발한 개발 언어의 신

- Go의 특징
	- 컴파일 언어
	- 시스템 프로그래밍을 위해 개발됨
	- C++과 같은 정적 타입 언어이며, Java/C#과 같이 가비지컬렉션이 있는 객체지향 언어임
	- 최소한의 키워드 만으로 개발이 가능하게 만듦

- Go의 사용사례
	- 시스템 수준의 애플리케이션
	- 웹 앱
	- 클라우드 네이티브 애플리케이션(도커, 쿠버네티스 등)
	- 분산 시스템
	- 데이터베이스 등

### Go를 배우기

#### Go Playground
- 웹사이트 제공 
	- https://go.dev/play/

- 로컬 컴퓨터에 설치
	- OS 플랫폼별 설치

### Go 설치
- https://go.dev/dl/ 에서 OS 플랫폼에 맞는 설치본 다운로드 및 섪치

	<img src="https://raw.githubusercontent.com/hugoMGSung/study-go/main/images/go001.png" width="500">

- 설치 시 명령 프롬프트를 관리자 권한으로 실행
	- 설치 위치의 go1.XX.Y.windows-amd64.msi 설치파일 실행
	- 삭제 시도 명령 프롬프트에서 위의 설치파일 실행 > remove 진행
	- 설치 후 윈도우 재시작

- Visual Studio Code 확장팩 설치
	- Go 설치
	- 명령 팔레트 > Go:Install/Update Tools 선택
	
	<img src="https://raw.githubusercontent.com/hugoMGSung/study-go/main/images/go002.png" width="600">

	- 출력 탭 확인, 모두 설치 후 [info] All tools successfully installed. You are ready to Go. :) 표시되면 완료


### GO 개발환경 시작해보기
- GOPATH 설정 경로에 bin, pkg, src 폴더를 생성
	- echo %GOPATH% 로 확인 -> %USERPROFILE%\go

- Visual Studio Code 실행
	- go-app 폴더 생성
	- main.go 생성
	- 터미널에서 명령어 실행
		```shell
		PS D:\01_Programming\100_HugoBank\Mine\study-go> go mod init go-app
		go: creating new go.mod: module go-app
		go: to add module requirements and sums:
				go mod tidy
		```
	- go.mod 파일 생성 확인
	- main.go 에 코드 작성
		```go
		package main

		import "fmt"

		func main() {
			name := "Go!"
			fmt.Println("VS Code for", name)
		}
		```
	- Ctrl + F5로 실행

		<img src="https://raw.githubusercontent.com/hugoMGSung/study-go/main/images/go003.png" width="730">

- 디버깅
	- 브레이크포인트 토글 후 F5 실행

		<img src="https://raw.githubusercontent.com/hugoMGSung/study-go/main/images/go004.png" width="730">

- 콘솔에서 실행
	- > go run main.go

- 실행파일 빌드
	- > go build main.go

### 기본 학습하기

#### 프로젝트에 관하여
- go mod init
	- go 모듈 초기화 명령
	- github 등에 모듈을 공유하기 위해서는 공식모듈명 사용 추천
		- 예) go mod init hugo83.com/goprog 

#### 소스코드 분석
- main.go
	- 1번 라인
		```go
		package main
		```

		- 실행가능한 앱임을 Go에게 알림. 모든 실행 파일에는 이 첫번째 라인이 있음

	- 3번 라인
		```go
		import "fmt"
		```

		- fmt 패키지를 여기에서 사용하겠음을 선언. 
		- fmt는 표준 라이브러리 [Go 패키지 설명](https://pkg.go.dev/fmt) 에 설명 발췌
			- Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
			- fmt 패키지는 C의 printf 및 scanf와 유사한 기능을 사용하여 형식화된 I/O를 구현함. '동사' 형식은 C에서 파생되었지만 더 간단함.

	- 5번 라인
		```go
		func main() {
		```

		- 메인 함수 선언. 보통 Entry point(시작점) 으로 불림. { 는 코드블럭의 시작을 알림

	- 6번 라인
		```go
			name := "Go!"
		```

		- 변수 name에 문자열 Go! 를 할당.

	- 	7번 라인
		```go
			fmt.Println("VS Code for", name)
		```

		- fmt.Println() - fmt 라이브러리에 포함된 Println() 함수를 호출
		- 여려 문자열 및 변수를 , 로 나누어 여러 번 사용 가능
		- fmt.Println() 은 os의 표준 출력 함수임

	- 8번 라인
		```go
			// 여기는 주석
		```

		- 한 줄 주석을 나타냄. 주석은 프로그램 실행에 영향을 주지 않음(당연!)

	- 9~10번 라인
		```go
			/* 여러줄 주석도
				   가능 함!! */
		```

		- 여러줄 주석도 사용할 수 있음.

	- 11번 라인
		```go
		}
		```

		- main 함수의 코드블럭 종료 모든 함수 선언이 완료됨.


