package linked_list

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// Two pointers, left and right, are set at the head node. Move the right pointer n steps forward. After doing that, **both pointers are exactly separated by n nodes apart**. Start moving both pointers forward until the right pointer reaches the last node.
func RemoveKthNodeFromEndOptimized(head *LinkedList, k int) *LinkedList {
	right, left := head, head

	for i := 0; i < k; i++ {
		// should also have a check that right.next is not nil, skipping for now
		right = right.Next
	}

	if right == nil {
		return head.Next
	}

	for right.Next != nil {
		right = right.Next
		left = left.Next
	}

	left.Next = left.Next.Next

	return head
}

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	total := 0

	tmp := head

	for tmp != nil {
		total++
		tmp = tmp.Next
	}

	target := total - k
	counter := 1
	current := head

	fmt.Println(false)

	if total == k {
		current.Value = current.Next.Value
		current.Next = current.Next.Next
		return
	}

	for current != nil {

		if counter == target {

			if current.Next != nil {
				current.Next = current.Next.Next
				current = current.Next
			} else {
				current.Next = nil
				current = current.Next
			}

			break
		}

		counter++
		current = current.Next
	}
}
