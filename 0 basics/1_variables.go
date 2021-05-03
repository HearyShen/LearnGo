package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

// run with command:
// go run src\1_variables.go

func main() {
	// 3 ways to declare and initialize a variable
	fmt.Printf("3 ways to declare and init a variable:\n")

	var a int = 100
	var b = "100"
	c := 0.17

	fmt.Printf("a value = %v (%T)\n", a, a)
	fmt.Printf("b value = %v (%T)\n", b, b)
	fmt.Printf("c value = %v (%T)\n", c, c)
	fmt.Println()

	// swap variables
	fmt.Printf("Easy way to swap variables:\n")
	v1 := 1
	v2 := 2
	fmt.Printf("before swap: v1 = %v (%T), v2 = %v (%T)\n", v1, v1, v2, v2)

	v1, v2 = v2, v1
	fmt.Printf("after swap: v1 = %v (%T), v2 = %v (%T)\n", v1, v1, v2, v2)
	fmt.Println()

	/*
		Integer
		signed: int8, int16, int32, int64
		unsigned: uint8, uint16, uint32, uint64
	*/
	var vUint16 uint16 = math.MaxUint8 + 1
	// vUint16 = math.MaxUint16 + 1		// src\1_variables.go:37:9: constant 256 overflows uint16
	fmt.Printf("vUint16 = %v (%T)\n", vUint16, vUint16)

	var vUint8 uint8 = uint8(vUint16)
	fmt.Printf("vUint8 = %v (%T)\n", vUint8, vUint8) // truncated: 0000 0001 (0000 0000)

	/*
		Float
		float32, float64
	*/
	var vFloat32 float32 = math.E
	var vFloat64 float64 = math.E
	fmt.Printf("vFloat32 = %f (%T)\n", vFloat32, vFloat32)
	fmt.Printf("vFloat64 = %.10f (%T)\n", vFloat64, vFloat64)

	/*
		Bool
		true/false, can not cast to integer types
	*/
	var vBool bool = true
	fmt.Printf("vBool = %v (%T)\n", vBool, vBool)

	/*
		String
	*/
	var vStr string = "你好, Go!"
	fmt.Printf("vStr = \"%s\" (%T)\n", vStr, vStr)
	fmt.Printf("byte len of vStr = %v\n", len(vStr))                    // 3*2 + 5*1 = 11, chinese character uses 3 bytes
	fmt.Printf("rune len of vStr = %v\n", utf8.RuneCountInString(vStr)) // 7
	// traverse each unicode character
	for i, h := range vStr {
		fmt.Printf("[%v:%c]", i, h)
	}
	fmt.Println()

	/*
		Pointer
	*/
	var ptrStr *string = &vStr
	fmt.Printf("ptrStr = %v (%T)\n", ptrStr, ptrStr)
	fmt.Printf("*ptrStr = %v (%T)\n", *ptrStr, *ptrStr)

	/*
		Struct
	*/
	var vStruct struct {
		id     int
		name   string
		salary float32
	}
	vStruct.id = 1
	vStruct.name = "Mike"
	vStruct.salary = 123.45
	fmt.Printf("vStruct = %v (%T)\n", vStruct, vStruct)
}
