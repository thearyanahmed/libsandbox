package linked_list

import (
	"fmt"
)

type Node struct {
	data int64
	next *Node
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
	node := Node{data: data, next: l.Head}

	l.Head = &node
	l.Len++
}

// TODO: update logic , check for if node exists in next.next
func (l *SinglyLinkedList) DeleteNthNode(i int) {
	if i > int(l.Len) {
		return
	}

	if i == 0 && (i+2) <= int(l.Len) {
		l.Head = l.Head.next.next
		return
	}

	targetNode := l.Head
	current := 0

	// 1 -> 2 -> 3 -> 4 -> 5
	// should be 1 -> 3 -> 4 -> 5
	for current < i {
		if targetNode.next != nil {
			targetNode = targetNode.next
			current++
		} else {
			break
		}
	}

	fmt.Println("target - 1 node => ", targetNode.data, current, i)
	// targetNode.next = targetNode.next.next
	// check for memory stuff
	targetNode.next = targetNode.next.next
}

func (l SinglyLinkedList) GetLength() int32 {
	return l.Len
}

func (l SinglyLinkedList) Print() {
	node := l.Head

	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}
