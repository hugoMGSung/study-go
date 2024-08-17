package main

import (
	"io"
	"io/ioutil"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	// 입력파일 열기
	fi, err := os.Open("./test.md")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	// 출력파일 생성
	fo, err := os.Create("./result.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buff := make([]byte, 1024)

	// 루프
	for {
		// 읽기
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		// 끝이면 루프 종료
		if cnt == 0 {
			break
		}

		// 쓰기
		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}

	//파일 읽기
	bytes, err := ioutil.ReadFile("./result.txt")
	if err != nil {
		panic(err)
	}
	//파일 쓰기
	err = ioutil.WriteFile("./next.log", bytes, 0)
	if err != nil {
		panic(err)
	}
}
