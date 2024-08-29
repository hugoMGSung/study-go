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

##### 진행 순서1
1. /init/main.go 기본 웹서버 실행 확인
2. /config/config.go 구성설정 작성
3. /init/cmd/cmd.go 작성
4. config.toml 에 서버 설정 구성
5. /init/main.go config 관련 내용 수정
6. /network/user.go 로 유저라우터 작성

	<img src="./images/img013.png" width="750">

	<img src="./images/img014.png" width="750">

7. /network/utils.go에 /network/root.go 에 있던 함수들 이전. Response관련 함수 추가
8. /network/user.go에 u.router.okResponse 관련 추가
9. /types/utils.go Reponse 구조체와 생성 함수 작성
10. /types/user.go JSON으로 리턴하는 User구조체와 HTTP메서드 별 구조체 생성

##### 진행 순서2
1. /service/root.go에 서비스 구조체와 이를 싱글톤으로 실행하는 변수 선언. 생성자 작성
2. /repository/root.go에 저장소 구초제, 싱글톤, 생성자 작성
3. 네트워크에서는 서비스만, 서비스에서는 리포지토리만 접근 하도록 레이아웃 분리 
4. /init/cmd/cmd.go Cmd 구조체에 리포지토리, 서비스 추가. NewCmd 생성내용에 추가
5. /service/user.go 생성
6. /repository/user.go 생성
7. /service/user.go CRUD 함수 작성
8. /repository/user.go CRUD 함수 작성
9. /network/user.go CRUD 함수에서 service 메서드 사용




