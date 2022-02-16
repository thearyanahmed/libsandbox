package linked_list

/**
 * Definition for singly-linked list.
 */
type listNodeA struct {
	Val  int
	Next *listNodeA
}

func removeNthFromEnd(head *listNodeA, n int) *listNodeA {

	if head == nil {
		return nil
	}

	if n < 1 {
		return head
	}

	length := 0

	node := head

	for node != nil {
		node = node.Next
		length++
	}

	if length < n {
		return head
	}

	if n == length {
		current := head
		head = head.Next
		current.Next = nil

		return head
	}

	node = head

	i := 0

	for node != nil {
		if i == length-n-1 {
			nodeTobeDeleted := node.Next
			node.Next = node.Next.Next
			nodeTobeDeleted.Next = nil
			break
		}
		i++
		node = node.Next
	}
	return head
}
