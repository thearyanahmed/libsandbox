package main

import (
	"fmt"
	"github.com/thearyanahmed/libsandbox/go/linked_list"
	_ "github.com/thearyanahmed/libsandbox/go/utils"
)

func main() {
	fmt.Println("hello")

	list := linked_list.NewSinglyLinkedList()

	for i := 0; i <= 10; i++ {
		list.Push(int64(i))
	}

	list.Print()
	fmt.Println(list.Mid().GetData())
}
