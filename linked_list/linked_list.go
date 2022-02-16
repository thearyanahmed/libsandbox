package linked_list

import (
	"fmt"
)

type Node struct {
	data    int64
	Next    *Node
	visited bool
}

type SinglyLinkedList struct {
	Head *Node
	Len  int32
}

func (node *Node) GetData() int64 {
	return node.data
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{Head: nil}
}

func (l *SinglyLinkedList) Push(data int64) {
	node := Node{data: data, Next: l.Head, visited: false}

	l.Head = &node
	l.Len++
}

// TODO: update logic , check for if node exists in Next.Next
func (l *SinglyLinkedList) DeleteNthNode(i int) {
	if i > int(l.Len) {
		return
	}

	if i == 0 && (i+2) <= int(l.Len) {
		l.Head = l.Head.Next.Next
		return
	}

	targetNode := l.Head
	current := 0

	// 1 -> 2 -> 3 -> 4 -> 5
	// should be 1 -> 3 -> 4 -> 5
	for current < i {
		if targetNode.Next != nil {
			targetNode = targetNode.Next
			current++
		} else {
			break
		}
	}

	// Check Next.Next and handle last element
	// targetNode.Next = targetNode.Next.Next
	// check for memory stuff
	targetNode.Next = targetNode.Next.Next

	l.Len--
}

func (l *SinglyLinkedList) NthFromLast(i int32) *Node {
	if i >= l.Len {
		return nil
	}

	nth := l.Len - i

	node := l.Head

	current := 0

	for int32(current) < nth {
		node = node.Next
		current++
	}

	return node
}

func (l *SinglyLinkedList) HasCycle() bool {
	node := l.Head

	for node != nil {
		if node.visited == true {
			return true
		}

		node.visited = true
		node = node.Next
	}

	return false
}

func (l SinglyLinkedList) GetLength() int32 {
	return l.Len
}

func (l SinglyLinkedList) Print() {
	node := l.Head

	for node != nil {
		fmt.Println(node.data)
		node = node.Next
	}
}
