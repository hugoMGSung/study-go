package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func init() {
	fmt.Println("수 타입 학습")
}

func main12() {
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

	// 정수형 최대값
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println(n1, n2, n3, n4)

	var n5 float32 = math.MaxFloat32
	var n6 float64 = math.MaxFloat64
	var n7 uint32 = math.MaxUint32 // + 1
	var n8 uint64 = math.MaxUint64 //+ 1

	fmt.Println(n5, n6, n7, n8)

	// 연산
	var n9 uint8 = 110
	var n10 uint8 = 90

	fmt.Println(n9+n10, n9-n10, n9*n10, n9/n10, n10%n9, n9<<2, n10>>2, ^n10)

	var n11 int32 = 1024
	var n12 float32 = 76.5

	//fmt.Println(n11 + n12) // 그냥 에러
	fmt.Println(float32(n11) + n12) // 또는
	fmt.Println(n11 + int32(n12))

	// 문자열을 정수형으로
	strInt := "100"
	intStr := string(strInt)
	fmt.Println(intStr, reflect.TypeOf(intStr)) // 100 string

	i, err := strconv.Atoi(strInt)
	fmt.Println(i, err, reflect.TypeOf(i)) //100 <nil> int

	strInt = "987654321"
	i8, err := strconv.ParseInt(strInt, 0, 8)
	i16, err := strconv.ParseInt(strInt, 0, 16)
	i32, err := strconv.ParseInt(strInt, 0, 32)
	i64, _ := strconv.ParseInt(strInt, 16, 64)

	fmt.Println(i8, err, reflect.TypeOf(i8))   // 127 <nil> int64
	fmt.Println(i16, err, reflect.TypeOf(i16)) // 32767 <nil> int64
	fmt.Println(i32, err, reflect.TypeOf(i32)) // 987654321 <nil> int64
	fmt.Println(i64, reflect.TypeOf(i64))      // 40926266145 <nil> int64
}
