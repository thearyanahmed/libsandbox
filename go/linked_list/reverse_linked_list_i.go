package linked_list

type ReverseLinkedListNode struct {
	Val  int
	Next *ReverseLinkedListNode
}

func ReverseLinkedList(head *ReverseLinkedListNode) *ReverseLinkedListNode {
	var prev *ReverseLinkedListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next

		head.Next = prev
		prev = head // current

	}
	return prev
}
