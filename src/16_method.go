package main

import "fmt"

// run with commands:
// go run src\16_method.go

/*
	In go, method is a function with recipient.
	Recipient can be any type, typically a struct, which means any type in
go can have its methods.

	Method can be defined in paradigm:

	func (recipient RecipientType) methodName(inputParams) (returnParams) {
		func body
	}
*/

type Student struct {
	Name string
	Age  uint8
	ID   uint64
}

// modify student's name with pointer to the instance
func (student *Student) setName(name string) {
	student.Name = name
}

// non-pointer, unable to modify the original instance
func (student Student) badSet(name string) {
	student.Name = name
}

func (student Student) print() {
	fmt.Printf("Student %s (ID: %v) is %v years old.\n",
		student.Name, student.ID, student.Age)
}

func main() {
	student1 := Student{
		Name: "Jack",
		Age:  12,
		ID:   1,
	}

	student1.print()

	student1.badSet("Little Jack")
	student1.print()

	student1.setName("Big Jack")
	student1.print()
}
