package main

import (
	"fmt"
	"time"
)

// run with commands:
// go run src\12_func.go

/*
	A standard paradigm of function in go:

	func func_name(input_params) (return_params) {
		func body
	}
*/

// a func with multiple inputs and single output
func add(x, y int) int {
	return x + y
}

// a func with multiple inputs and multiple outputs
// return values can be named
func div(dividend, divisor int) (quotient, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}

func addMul(x, y int) (int, int) {
	vAdd := x + y
	vMul := x * y
	return vAdd, vMul
}

// a function with input but no return value
func echo(s string) {
	fmt.Println(s)
}

// input a func as a callback func
func traverse(arr []int, handler func(num int)) {
	for _, v := range arr {
		handler(v)
	}
}

// pass a pointer to func
func increase(x *int) {
	*x = 2
}

func main() {
	// call a named function
	fmt.Println(add(1, 2))
	fmt.Println(div(8, 5))
	fmt.Println(addMul(2, 3))
	echo("Hello, world!")

	// define and call an anonymous func
	vMul := func(x, y int) int {
		return x - y
	}(1, 2)
	fmt.Println(vMul)

	curTime := func() {
		fmt.Println(time.Now())
	}
	curTime()

	// use an anonymous func as callback function
	arr := []int{1, 2, 3, 4, 5}
	traverse(arr, func(num int) {
		fmt.Print(num*num, " ")
	})
	fmt.Println()

	// pass num to func by its pointer
	vNum := 1
	fmt.Println("Before: ", vNum)
	increase(&vNum)
	fmt.Println("After: ", vNum)
}
