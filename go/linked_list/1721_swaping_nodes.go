package linked_list

type listNodeB struct {
	Val  int
	Next *listNodeB
}

func swapNodes(head *listNodeB, k int) *listNodeB {
	arr := []int{}

	node := head

	length := 0

	for node != nil {
		arr = append(arr, node.Val)
		node = node.Next

		length++
	}

	// swap
	// length = 5

	arr[k-1], arr[length-k] = arr[length-k], arr[k-1]

	h := &listNodeB{Val: arr[0], Next: nil}
	j := h

	//fmt.Println("array",arr)

	for i := 1; i < length; i++ {
		node := listNodeB{Val: arr[i], Next: nil}

		h.Next = &node
		h = h.Next
	}

	return j

}
