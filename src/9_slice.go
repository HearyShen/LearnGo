package main

import "fmt"

// run with commands:
// go run src\9_slice.go

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	printSlice("slice", slice)
	printSlice("subSlice", slice[3:])

	slice = append(slice, 7, 8)
	printSlice("slice", slice)
	printSlice("subSlice", slice[3:])
}

func printSlice(tag string, slice []int) {
	fmt.Printf("%s: date = %v, len = %d, cap = %d, addr = %p\n", tag, slice, len(slice), cap(slice), &slice)
}
