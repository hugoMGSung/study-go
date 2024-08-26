package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func init() {
	fmt.Println("패키지 문서화")
}

func main81() {
	// 외부 저장소 설치
	// 1. import 선언 후 폴더 이동 후 go get 설치
	// 2. go get 패키지 주소로 설치
	// go get github.com/tealeg/xlsx
	file := "sample.xlsx"

	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		panic("Excel읽이 오류!!")
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\t", text) // 각 셀 데이터 출력
			}
			fmt.Println()
		}
	}

	// 쓰기
	nfile := xlsx.NewFile()

	sheet, err := nfile.AddSheet("Sheet1")
	if err != nil {
		panic(err)
	}

	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "VALUE"

	err = nfile.Save("./lastTest.xlsx")
	if err != nil {
		panic(err)
	}
}
