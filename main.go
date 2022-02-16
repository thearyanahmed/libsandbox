package main

import (
	"fmt"

	"github.com/thearyanahmed/rknot/linked_list"
)

func main() {
	fmt.Println("program start")

	l := linked_list.NewSinglyLinkedList()

	l.Push(10)
	l.Push(20)
	l.Push(30)

	l.Print()
}
