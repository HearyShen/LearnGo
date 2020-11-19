package main

import (
	"flag"
	"fmt"
)

// run with command:
// go run src\2_flag.go --username mike --password 123456 --id 1

func main() {
	// argument name, default value, tips
	var username *string = flag.String("username", "guest", "a string of username")

	var password string
	flag.StringVar(&password, "password", "123", "a string of password")

	id := flag.Int("id", 0, "an integer of id")

	flag.Parse()
	fmt.Printf("username = %v, password = %v, id = %v\n", *username, password, *id)
}
