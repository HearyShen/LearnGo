package main

import "fmt"

// run with commands:
// go run src\14_struct.go

/*
	struct can be defined in paradigm:

	type structName struct {
		value1 valueType1
		value2 valueType2
		...
	}
*/

type Person struct {
	Name  string
	Birth string
	ID    uint64
}

func main() {
	// declare a struct variable
	var person1 Person
	person1.Name = "Mike"
	person1.Birth = "1990-1-2"
	person1.ID = 1
	fmt.Println(person1)

	// new a struct variable
	person2 := new(Person) // person2 is a pointer
	person2.Name = "Tom"
	person2.Birth = "1991-2-3"
	person2.ID = 2
	fmt.Println(person2)

	// another way to new a person with empty initial values
	person3 := &Person{} // person3 is a pointer
	person3.Name = "Nancy"
	person3.Birth = "1992-3-4"
	person3.ID = 3
	fmt.Println(person3)

	// create an object with initial values
	person4 := Person{
		Name:  "Jack",
		Birth: "1993-4-5",
		ID:    4,
	}
	fmt.Println(person4)

	person5 := &Person{ // person5 is a pointer
		Name:  "John",
		Birth: "1994-5-6",
		ID:    5,
	}
	fmt.Println(person5)
}
