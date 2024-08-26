package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("파일 입출력3")
}

func main68() {
	file, err := os.Open("test.txt")
	errCheck1(err)

	defer file.Close() // 이건 위에 있어도 됨

	fileInfo, err := file.Stat() //파일 사이즈 확인 위해 정보 획득
	errCheck2(err)
	fd1 := make([]byte, fileInfo.Size()) //슬라이스에 읽은 내용 담는다.
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보", fileInfo, "\n")
	fmt.Println("파일 이름", fileInfo.Name(), "\n")
	fmt.Println("파일 크기", fileInfo.Size(), "\n")
	fmt.Println("파일 수정시간", fileInfo.ModTime(), "\n")
	fmt.Printf("읽기 작업 완료 (%d bytes)\n\n", ct1)
	//fmt.Println(fd1) //타입 변환 없을 경우
	fmt.Println(string(fd1)) //타입 변환 한 경우

	// //읽기 예제2(탐색 : Seek(Offset, whence))
	o1, err := file.Seek(20, 0) // whence - 0: 처음위치 , 1:  현재 위치 , 2: 마지막 위치
	errCheck1(err)

	fd2 := make([]byte, 20)
	ct2, err := file.Read(fd2)
	errCheck1(err)

	fmt.Printf("읽기 작업 완료 (%d bytes) (%d ret)\n\n", ct2, o1)
	fmt.Println(string(fd2), "\n\n")

	//읽기 예제3
	o2, err := file.Seek(0, 0)
	errCheck1(err)
	fd3 := make([]byte, 50)         // 이 크기만 사용
	ct3, err := file.ReadAt(fd3, 8) //offset 위치부터 읽어온다.
	errCheck1(err)

	fmt.Printf("읽기 작업 완료 (%d bytes) (%d ret)\n\n", ct3, o2)
	fmt.Println(string(fd3), "\n")
}
