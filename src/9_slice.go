package main

import "fmt"

// run with commands:
// go run src\9_slice.go

// slice is a variable length sequence of data

func main() {
	// [...]int{} marks a fixed length array,
	// while []int{} declares a variable length slice
	slice := []int{1, 2, 3, 4, 5, 6}
	printSlice("slice", slice)
	printSlice("subSlice", slice[3:])

	// slice can be appended with one or more elements
	// If it has sufficient capacity, the destination is resliced to accommodate the new elements.
	// If it does not, a new underlying array will be allocated.
	slice = append(slice, 7, 8)
	printSlice("slice", slice)
	printSlice("subSlice", slice[3:])

	// slice can also be created from an array
	arr := [...]int{11, 22, 33, 44, 55}
	arrSlice := arr[:3]
	printSlice("arrSlice", arrSlice)

	// slice can also be dynamically created with make([]T, size, cap)
	madeSlice := make([]int, 8, 16)
	printSlice("makeSlice", madeSlice)

	// slice can be copied from src slice to dest slice
	copy(madeSlice, slice)
	printSlice("copiedSlice", madeSlice)
}

func printSlice(tag string, slice []int) {
	fmt.Printf("%s: date = %v, len = %d, cap = %d, addr = %p\n", tag, slice, len(slice), cap(slice), &slice)
}
