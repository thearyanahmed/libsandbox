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
	l.Push(40)
	l.Push(50)

	last := l.NthFromLast(2)

	if last != nil {
		fmt.Println(last.GetData())
	}

	l.Print()
}
