package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func init() {
	fmt.Println("문자열 타입 학습")
}

func main13() {
	var str1 string = "Hello, Go"
	var str2 string = "안녕"
	//
	var str3 string = "\uc548\uB155"         // 한글 유니코드(4자리는 소문자 -> u)
	var str4 string = "\U0000C548\U0000B155" // 한글 유니코드(8자리는 소문자 -> U)
	// https://cloudedi.tistory.com/entry/UTF-8-UTF-16-UTF-32-%ED%95%9C%EA%B8%80-%EC%9D%B8%EC%BD%94%EB%94%A9
	var str5 string = "\xec\x95\x88\xeb\x85\x95" // UTF-8 인코딩 바이트 코드
	// https://www.browserling.com/tools/utf8-encode

	fmt.Println(str1, str2, str3, str4, str5)

	// 길이
	fmt.Println(len(str1), len(str2), utf8.RuneCountInString(str3), len([]rune(str4)))

	var str6 string = "Mie, Golang"
	var str7 string = "미고랭"

	// 인덱싱
	fmt.Println(str6[3], str7[1])
	fmt.Printf("%c %c\n", str6[3], str7[1])

	conStr := []rune(str7)
	fmt.Printf("%c %c\n", str6[2], conStr[1])

	for i, char := range str6 {
		fmt.Printf("%c[%d]", char, i)
	}
	fmt.Println()

	for i, char := range []rune(str7) {
		fmt.Printf("%c[%d]", char, i)
	}
	fmt.Println()

	// 문자열 슬라이스 - 파이썬과 유사
	fmt.Println(str6[0:2], str6[:2], str6[7:])
	fmt.Println(str7[0:3], str7[:3], str7[6:])
	fmt.Println(([]rune(str7))[0:2])
	fmt.Printf("%c\n", ([]rune(str7))[0:2])

	// 비교연산
	fmt.Println(str2 == str4)
	fmt.Println(str6 < str7)

	// 이하는 strings 패키지 사용
	// 백택(`) 은 여러줄에 편리해도 탭, 공백 등의 문제가 존재
	str10 := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis eget scelerisque velit. 
			  Aliquam erat volutpat. Ut eget placerat turpis. Quisque euismod suscipit cursus. `
	str11 := "Pellentesque dignissim, enim quis efficitur aliquet, ipsum mi ullamcorper lacus, bibendum bibendum nibh urna quis libero. " +
		"Integer gravida quam quis enim iaculis luctus. Nullam et auctor dui."
	fmt.Println(str10 + str11)

	strSet := []string{} // 선언
	strSet = append(strSet, str10)
	strSet = append(strSet, str11)
	fmt.Println(strSet)
	// fmt.Println(strings.Join(strSet, " "))
	fmt.Println(strings.Replace(strings.Join(strSet, " "), "\t", "", -1))
}
