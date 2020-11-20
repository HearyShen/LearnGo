package main

import "fmt"

// run with commands:
// go run src\7_forloop.go

// Golang only has for loop
// Golang does not support while and do-while loop

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
