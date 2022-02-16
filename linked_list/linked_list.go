package linked_list

import "fmt"

type Node struct {
	data int64
	next *Node
}

type SinglyLinkedList struct {
	Head *Node
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
}

func (l SinglyLinkedList) Print() {
	node := l.Head

	fmt.Println("head", node.data)

	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}
