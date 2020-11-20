package main

import (
	"flag"
	"fmt"
)

// run with command:
// go run src\5_ifelse.go --score 100

func main() {
	score := flag.Int("score", -1, "Score of a test")

	flag.Parse()
	fmt.Printf("score = %v (%T)\n", *score, *score)

	if *score < 60 {
		fmt.Println("Fail to pass")
	} else if *score < 80 {
		fmt.Println("Fine")
	} else if *score <= 100 {
		fmt.Println("Excellent")
	} else {
		fmt.Println("Wrong score")
	}
}
