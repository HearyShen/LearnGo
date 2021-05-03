package main

import "fmt"

// run with commands:
// go run src\16_interface2.go

/*
	struct can embed anonymous attributes (type-only) to implement composition relation.

	standard embedded struct type paradigm:

	type A struct {
		typeB
		typeC
	}
*/

type Swimming struct {
}

func (swim *Swimming) swim() {
	fmt.Println("swimming")
}

type Flying struct {
}

func (flying *Flying) fly() {
	fmt.Println("flying")
}

// Wild Duck can swim and fly
type WildDuck struct {
	Swimming
	Flying
}

// Domestic Duck can only swim
type DomesticDuck struct {
	Swimming
}

func main() {
	wildDuck := WildDuck{}
	wildDuck.fly()
	wildDuck.swim()

	domesticDuck := DomesticDuck{}
	domesticDuck.swim()
}
