package linked_list

type (
	ListNode struct {
		Val  int
		Next *ListNode
	}
)

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	node := list1
	var firstHalf, lastHalf *ListNode
	for i := 0; i <= b; i++ {
		if i == a-1 {
			firstHalf = node
		}
		if i == b {
			lastHalf = node
		}
		node = node.Next
	}
	firstHalf.Next = list2
	
	node = list2
	for node.Next != nil {
		node = node.Next
	}
	node.Next = lastHalf.Next
	return list1
}
