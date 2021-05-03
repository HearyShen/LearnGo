package main

import "fmt"

// run with command:
// go run src\4_type.go

type aliasInt = int // declare an alias for int
type myInt int      // defines a new type

// Person is type of a struct
type BasicPerson struct {
	name string
	age  uint8
}

func main() {
	var vAliasInt aliasInt
	fmt.Printf("vAliasInt = %v (%T)\n", vAliasInt, vAliasInt)

	var vMyInt myInt
	fmt.Printf("vMyint = %v (%T)\n", vMyInt, vMyInt)

	var person BasicPerson
	person.name = "Mike"
	person.age = 20
	fmt.Printf("person = %v (%T)\n", person, person)
}
