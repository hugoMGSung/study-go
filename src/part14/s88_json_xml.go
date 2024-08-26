package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Hero struct {
	Name   string
	Age    int
	Active bool
}

type Heros struct {
	Hero []Hero
}

func main88() {
	// Go 데이타
	hero := Hero{"IronMan", 50, false}

	// JSON 인코딩
	jsonBytes, err := json.Marshal(hero)
	if err != nil {
		panic(err)
	}

	// JSON 바이트를 문자열로 변경
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	var dhero Hero
	err1 := json.Unmarshal(jsonBytes, &dhero)
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(dhero.Name, dhero.Active)

	// XML 처리
	hero2 := Hero{"Captain America", 90, true}

	xmlBytes, err := xml.Marshal(hero2)
	if err != nil {
		panic(err)
	}

	xmlString := string(xmlBytes)
	fmt.Println(xmlString)

	// XML 읽기
	file, err := os.Open("./Heros.xml")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	// XML 디코딩
	var heros Heros
	xmlerr := xml.Unmarshal(data, &heros)
	if xmlerr != nil {
		panic(xmlerr)
	}

	fmt.Println(heros)
}
