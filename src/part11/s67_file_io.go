package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func init() {
	fmt.Println("파일 입출력2")
}

func main67() {
	// csv 파일 쓰기
	// excel의 경우 https://github.com/tealeg/xlsx
	// bufio - 파일 용량이 클경우 버퍼 사용 권장
	file, err := os.Create("first.csv")
	errCheck1(err)

	defer file.Close()

	// csv writer
	// wr := csv.NewWriter(file)
	wr := csv.NewWriter(bufio.NewWriter(file))

	// CSV쓰기
	wr.Write([]string{"Hero", "Power"})

	wr.Write([]string{"IronMan", "5.0"})
	wr.Write([]string{"AntMan", "2.5"})
	wr.Write([]string{"Hulk", "4.0"})
	wr.Write([]string{"Thor", "4.8"})
	wr.Write([]string{"Captain America", "3.9"})

	wr.Flush() // 버퍼 사용후 비우기

	fi, err := file.Stat()
	errCheck1(err)
	fmt.Printf("CSV 쓰기 후 파일 크기(%d bytes)\n", fi.Size())
	fmt.Println("CSV 파일명", fi.Name())

}
