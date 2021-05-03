package main

import "fmt"

// run with commands:
// go run src\16_interface2.go

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

type Printer interface {
	Print(interface{})
}

type FuncCaller func(p interface{})

func (funcCaller FuncCaller) Print(p interface{}) {
	funcCaller(p)
}

func main() {
	// Printer is the abstraction of printer
	// FuncCaller func is the implementation of printer
	// printer can call Printer.Print implemented by FuncCaller's Print
	var printer Printer
	printer = FuncCaller(func(p interface{}) {
		fmt.Println(p)
	})
	// cast an anonymous function to FuncCaller type
	// then printer calls Print implemented by FuncCaller
	printer.Print("Golang is Good!")
}
