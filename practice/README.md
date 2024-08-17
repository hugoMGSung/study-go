## GO 예제

### 파일 입출력
- os.Open() 파일 열때
- os.Create() 새 파일 생성
- Read(), Write() 로 파일 읽고 쓰기
- 파일이 크지 않으면 ReadFile(), WriteFile()

### HTTP GET 호출(헤더포함)
- Go 1.16 부터 io/ioutil 패키지가 deprecated ioutil.ReadAll -> io.ReadAll() 로 변경

### HTTP POST 호출
- 