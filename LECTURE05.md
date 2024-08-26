## Go언어 베이직

### Go언어 실습

#### 파일 입출력
- Create : 새 파일 작성 및 파일열기
- Close : 리소스 닫기
- Write, WriteString, WriteAt 
- Read, ReadAt
- OS별 권한 주의
- 에러처리 같이!!
- file.Stat() 파일 상태

- ioutil을 사용했지만, 1.16부터 deprecated. 사용하지 말 것

#### 내부 패키지 제작 문서화/설치
- 패키지를 잘 만들어서 제대로 전달해야 함
- 문서화 중요!!

- go install
- godoc -http=:6660

<img src="./images/img006.png" width="600">

<img src="./images/img007.png" width="730">


#### 외부 저장소 패키지 설치 및 사용
- 외부 패키지
	- gorilla/websocket
	- 
- GOPATH 아래 pkg에 설치

- 깃허브에서 검색후 프로토콜 외 복사 후 import 추가 
	- github.com/tealeg/xlsx

#### HTTP 처리
- HTTP GET
- HTTP POST

#### JSON / XML

#### DB연동
-  표준패키지 database/sql을 사용
- 

#### 간단 크롤링 제작
- 기본 크롤링 제작
	- MySQL: https://github.com/go-sql-driver/mysql
	- MSSQL: https://github.com/denisenkom/go-mssqldb
	- Oracle: https://github.com/rana/ora
	- Postgres: https://github.com/lib/pq
	- SQLite: https://github.com/mattn/go-sqlite3
	- DB2: https://bitbucket.org/phiggins/db2cli

- MySQL 사용
	- go get github.com/go-sql-driver/mysql

#### 기타
- 추천 사이트
	- https://go.dev/doc/
	- 고를 배워야 하는 이유 - https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65
	- https://www.slant.co/topics/1412/~best-web-frameworks-for-go
	- 웹 프레임워크 - https://revel.github.io/
	- 커뮤니티 - https://golangkorea.github.io/post/go-start/feature/

