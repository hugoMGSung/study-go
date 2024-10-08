package main

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	fmt.Println("파일 입출력3")
}

func main70() {
	file, err := os.OpenFile("test2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))

	//bufio 파일 쓰기 예제
	wt := bufio.NewWriter(file) //Writer 반환(버퍼 사용)
	wt.WriteString("Mie Golang!\n File Write Test1!\n")
	wt.Write([]byte("Mie Golang!\n File Write Test2!\n"))

	// 버퍼정보
	fmt.Printf("사용한 Buffer Size (%d bytes)\n", wt.Buffered())
	fmt.Printf("남은 Buffer Size (%d bytes)\n", wt.Available())
	fmt.Printf("전체 Buffer Size (%d bytes)\n", wt.Size())

	wt.Flush() //버퍼 비우고 반영(버퍼에 내용을 디스크에 기록) 없으면 안됨!!!!

	//에러 체크
	errCheck1(err)

	fmt.Println("쓰기 작업 완료\n\n")

	rt := bufio.NewReader(file) //Reader 반환
	fi, err := file.Stat()
	errCheck1(err)

	b := make([]byte, fi.Size())

	fmt.Println("파일 정보 출력 : ", fi)
	fmt.Println("파일 이름 : ", fi.Name())
	fmt.Println("파일 크기 : ", fi.Size())
	fmt.Println("파일 수정 시간 : ", fi.ModTime())

	file.Seek(0, 0)
	data, _ := rt.Read(b) //읽기(ReadLine, ReadByte, ReadBytes 등)

	fmt.Printf("전체 Buffer Size : (%d bytes)\n", rt.Size())
	fmt.Printf("읽기 작업 완료 : (%d bytes)\n", data)
	fmt.Println("=============================================")
	fmt.Println(string(b))

	defer file.Close() // 어디쓰나 관계없음
}
