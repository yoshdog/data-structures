package main

import "github.com/yoshdog/data-structures/pkg/datastructure"

func main() {
	testLinkedList := &datastructure.LinkedList[string]{}

	testLinkedList.Prepend("a")
	testLinkedList.Prepend("b")
	testLinkedList.Prepend("c")
	testLinkedList.Prepend("d")

	testLinkedList.Delete("d")
	testLinkedList.Delete("b")
}
