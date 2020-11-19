package main

import "fmt"

// run with command:
// go run src\3_const.go

func main() {
	const helloStr = "Hello, world!"
	const (
		name   = "Mike"
		salary = 123.45
	)
	fmt.Printf("%v (%T)", name, name)

	name = "Tom" // cannot assign to name
	fmt.Printf("%v (%T)", name, name)

	ptrName := &name // cannot take the address of name
	*ptrName = "Tom"
	fmt.Printf("%v (%T)", name, name)
}
