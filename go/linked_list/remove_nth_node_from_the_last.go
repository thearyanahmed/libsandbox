package linked_list

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	// Write your code here.
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
