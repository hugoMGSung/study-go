## Go언어 베이직

### Go언어 실습

#### GoLang Backend Development

##### 서버구조
- root
	- config - 서버 설정, env -> toml 사용
	- init - 서버 시작점 설정
		- cmd - 
	- network - 네트워크 api 설정 등
	- repository - 세션 정의 등 db처리
	- service - network와 repository 다리 역할
	- types - 여러 타입 선언
		- errors 

##### Go 초기화
- 명령프롬프트 > Go: Initialize go.mod 선택
	- backend-server 로 입력
- /init/cmd.go 테스트

##### config
- 패키지 설치
	- go get github.com/naoina/toml

##### gin 프레임워크
- 패키지 설치
	- go get github.com/gin-gonic/gin

##### 진행 순서
1. /init/main.go 기본 웹서버 실행 확인
2. /config/config.go 구성설정 작성
3. /init/cmd/cmd.go 작성
4. config.toml 에 서버 설정 구성
5. /init/main.go config 관련 내용 수정
6. /network/user.go 로 유저라우터 작성

	<img src="./images/img013.png" width="750">

	<img src="./images/img014.png" width="750">

7. next





