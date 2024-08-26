package main

import (
	"fmt"
	//"io/ioutil" // "io/ioutil" is deprecated: As of Go 1.16
	"os"
)

func init() {
	fmt.Println("파일 입출력3")
}

func main69() {
	//파일 읽기, 쓰기 -> ioutil 패키지 활용
	//더욱 편리하고 직관적으로 파일을 읽고 쓰기 가능
	//아래 메소드 확인 가능
	//WriteFile(), ReadFile(), ReadAlll() 등 사용 가능
	//https://golang.org/pkg/io/ioutil/

	s := "Hello Golang!\n File Write Test!\n"

	//파일모드(chmod, chown, chgrp) -> 퍼미션
	//읽기 : 4 , 쓰기 : 2 , 실행 : 1
	//소유자, 그룹, 기타 사용자 순서 (644)
	//https://golang.org/pkg/os/#FileMode
	// ioutil.WriteFile is deprecated: As of Go 1.16, this function simply calls [os.WriteFile]
	//err := ioutil.WriteFile("test_write1.txt", []byte(s), os.FileMode(0644))
	err := os.WriteFile("test_write1.txt", []byte(s), os.FileMode(0644))
	errCheck1(err)

	// ioutil.ReadFile is deprecated: As of Go 1.16, this function simply calls [os.ReadFile].
	//data, err := ioutil.ReadFile("test.txt")
	data, err := os.ReadFile("test.txt")
	errCheck1(err)

	fmt.Println("=====================================")
	fmt.Println(string(data))
	fmt.Println("=====================================")
}
