package main

import (
	"fmt"
	"reflect"
)

func init() {
	fmt.Println("인터페이스 심화")
}

func diagnotics(v interface{}) {
	fmt.Printf("타입 : (%T) ", v)
	fmt.Println(v)
}

func type_check(arg interface{}) {
	switch arg.(type) {
	case int, int8, int32, int64:
		fmt.Println("Int type >", arg)
	case uint, uint8, uint32, uint64:
		fmt.Println("Unsigned Int type >", arg)
	case float32, float64:
		fmt.Println("Float type >", arg)
	case string:
		fmt.Println("String type >", arg)
	case bool:
		fmt.Println("Boolean type >", arg)
	case nil:
		fmt.Println("Nil isn't empty! >", arg)
	default:
		fmt.Println("Undefined type >", arg)
	}
}

func main29() {
	// 동적타입 (타언어의 Object)
	var varz interface{}
	diagnotics(varz) // nil
	varz = 10
	diagnotics(varz)
	varz = 3.141592
	diagnotics(varz)
	varz = "String"
	diagnotics(varz)
	varz = byte(64)
	diagnotics(varz)
	varz = true
	diagnotics(varz)

	// 형변환 제약
	var origin interface{} = 100
	change1 := origin
	change2 := origin.(int) // 원래 타입으로만 변환가능
	// change3 := origin.(float64) // error
	fmt.Println(origin, reflect.TypeOf(origin))
	fmt.Println(change1, reflect.TypeOf(change1))
	fmt.Println(change2, reflect.TypeOf(change2))

	// 타입 검사
	if v, res := origin.(int); res {
		fmt.Println(res, v)
	}

	// 타입 검사 2
	type_check(10)
	type_check(3.141592)
	type_check("String")
	type_check(byte(64))
	type_check(false)
	type_check(nil)
}
