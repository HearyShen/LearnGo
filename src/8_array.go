package main

import "fmt"

func main() {
	// init array
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)

	// init array with initial values
	languages := [...]string{"C", "C++", "Java"}
	fmt.Println(languages)

	// init array with new(type), returns a pointer
	nations := new([3]string)
	nations[0] = "China"
	nations[1] = "India"
	nations[2] = "Japan"
	fmt.Println(*nations)
}
