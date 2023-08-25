package main

import (
	"fmt"
	_ "github.com/thearyanahmed/libsandbox/go/utils"
)

type linkedList struct {
	Val  int
	Next *linkedList
}

func (l *linkedList) insert(a int) {
	c := l

	for {
		if c.Next == nil {
			c.Next = &linkedList{Val: a}
			break
		}

		c = c.Next
	}
}

func main() {
	fmt.Println("hello")

	arr := []int{1, 2, 3, 4, 5, 6, 8}

	list := linkedList{
		Val:  arr[0],
		Next: nil,
	}

	for i := 1; i < len(arr); i++ {
		list.insert(arr[i])
	}

	node := &list

	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
