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