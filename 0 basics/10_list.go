package main

import (
	"container/list"
	"fmt"
)

// run with commands:
// go run src\10_list.go

// list.List in go is a doubly linked list

func main() {
	// create a list
	numsList := list.New()

	// append elements to list with PushBack
	for i := 1; i < 10; i++ {
		numsList.PushBack(i)
	}
	printList(numsList)

	// add elements to front with PushFront
	first := numsList.PushFront(0)
	printList(numsList)

	// elements can be removed
	numsList.Remove(first)
	printList(numsList)
}

func printList(srcList *list.List) {
	for node := srcList.Front(); node != nil; node = node.Next() {
		fmt.Print(node.Value, " ")
	}
	fmt.Println()
}
