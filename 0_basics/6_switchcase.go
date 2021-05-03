package main

import (
	"flag"
	"fmt"
)

// run with command:
// go run src\6_switchcase.go -score 100 -course CPP

func main() {
	score := flag.Int("score", -1, "Score of a test")
	course := flag.String("course", "CPP", "Course name")
	flag.Parse()

	fmt.Printf("score = %v (%T)\n", *score, *score)
	fmt.Printf("course = %v (%T)\n", *course, *course)

	switch {
	case *score < 60:
		fmt.Println("fail to pass")
		break
	case *score < 80:
		fmt.Println("fine")
		break
	case *score <= 100:
		fmt.Println("excellent")
		break
	default:
		fmt.Println("wrong score")
	}

	switch *course {
	case "CPP":
		fmt.Println("C plus plus")
		break
	case "PY":
		fmt.Println("Python")
		break
	default:
		fmt.Println("Unknown")
	}
}
