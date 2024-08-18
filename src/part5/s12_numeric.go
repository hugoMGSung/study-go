package main

import (
	"fmt"
)

func init() {
	fmt.Println("수 타입 학습")
}

func main() {
	// 정수, 실수, 복소수
	// 32bit, 64bit, unsigned(양수)
	// 정수 : 8진수(0), 16진수(0x), 10진수
	var num1 int = 10
	var num2 int = -10
	var num3 int = 012  // 8진수
	var num4 int = 0x0A // 16진수

	fmt.Println(num1, num2, num3, num4)

	// 정수의 문자화
	var num5 byte = 70
	// %d, %c, %o, %x
	fmt.Printf("%d, %c\n", num5, num5)
	var num6 rune = 0xD540 // 유니코드
	fmt.Printf("%d, %c\n", num6, num6)

	// 소수점 사용(7자리)
	var num7 float32 = 44.1378373
	// 지수 표기법(15자리)
	var num8 float64 = 7.786325e-10
	fmt.Println(num7, num8)
	// float64(10.0 - 0.1) - 부동소수점 오차

	// 복소수 형(complex number)
	// complex64(32bit 실수 + 허수)
	// complex128(64bit 실수 + 허수)
	var num9 complex64 = 1 + 2i
	num10 := 3 + 4i
	num11 := complex(5, 6) //complex128
	var num12 complex128 = 7 + 8i
	num13 := complex64(9 + 10i)

	fmt.Println(num9, num10, num11, num12, num13)

	// real() : 실수부 출력 / imag() : 허수부 출력
	fmt.Println(real(num9), imag(num12))
}
