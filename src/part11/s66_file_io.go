package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("파일 입출력1")
}

func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

func errCheck2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main66() {
	// 파일 쓰기
	file, err := os.Create("./first.txt")
	errCheck1(err)
	// errCheck2(err)

	// 리소스 해제
	defer file.Close()

	// 쓰기
	s1 := []byte{77, 105, 101, 44, 71, 111, 108, 97, 110, 103}
	n1, err := file.Write([]byte(s1)) // 문자열 -> byte 슬라이스 형으로 변환 후 쓰기
	errCheck2(err)
	fmt.Println("쓰기완료", n1)
	file.Sync() // Write Commit(없어도 됌)

	// 쓰기2
	s2 := "Lorem ipsum dolor sit amet,\n consectetur adipiscing elit,\n"
	n2, err := file.WriteString(s2)
	errCheck2(err)
	fmt.Println("쓰기완료", n2)
	file.Sync()

	s3 := "Mie, Golang!"
	n3, err := file.WriteAt([]byte(s3), 40) // space 40번째 자리에
	errCheck2(err)
	fmt.Println("쓰기완료", n3)
	file.Sync()
}
