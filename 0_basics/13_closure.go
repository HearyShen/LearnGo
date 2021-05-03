package main

import "fmt"

// run with commands:
// go run src\13_closure.go

// closure is a function carrying state

func createCounter(initial int) func() int {
	if initial < 0 {
		initial = 0
	}

	return func() int {
		initial++
		return initial
	}
}

func main() {
	counter1 := createCounter(0)
	fmt.Println(counter1()) // 1
	fmt.Println(counter1()) // 2

	counter2 := createCounter(10)
	fmt.Println(counter2()) // 11

	fmt.Println(counter1()) // 3
}
