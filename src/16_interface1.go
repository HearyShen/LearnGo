package main

import "fmt"

// run with commands:
// go run src\16_interface1.go

/*
	standard interface paradigm:

	type interfaceName interface {
		func1(inputParams) (returnParams)
		func2(inputParams) (returnParams)
		func3(inputParams) (returnParams)
	}

	If the interfaceName is in uppercase, its a public interface.
	If the function name is in uppercase, its a public function.
	A public function can be accessed outside of the package,
	otherwise, non-public function can only be accessed inside of
	the package.
*/

type Cat interface {
	CatchMouse()
}

type Dog interface {
	Bark()
}

type CatDog struct {
	Name string
}

// type CatDog implements functions in interface Cat
func (catDog *CatDog) CatchMouse() {
	fmt.Printf("%s is catching mice!\n", catDog.Name)
}

// type CatDog implements functions in interface Dog
func (catDog *CatDog) Bark() {
	fmt.Printf("%s is barking!\n", catDog.Name)
}

func main() {
	// catDog is a pointer to CatDog instance
	catDog := &CatDog{
		Name: "Tom",
	}

	// declare Cat interface and point to CatDog type
	var cat Cat
	cat = catDog
	cat.CatchMouse()

	// declare Dog interface and point to CatDog type
	var dog Dog
	dog = catDog
	dog.Bark()
}
