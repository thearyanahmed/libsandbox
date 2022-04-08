package linked_list

type LRUCacheWorking struct {
	head, tail *NodeA
	Keys       map[int]*NodeA
	Cap        int
}

type NodeA struct {
	Key, Val   int
	Prev, Next *NodeA
}

func Constructor(capacity int) LRUCacheWorking {
	return LRUCacheWorking{Keys: make(map[int]*NodeA), Cap: capacity}
}

func (this *LRUCacheWorking) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCacheWorking) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node = &NodeA{Key: key, Val: value}
		this.Keys[key] = node
		this.Add(node)
	}
	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}
}

func (this *LRUCacheWorking) Add(node *NodeA) {
	node.Prev = nil
	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}

func (this *LRUCacheWorking) Remove(node *NodeA) {
	if node == this.head {
		this.head = node.Next
		node.Next = nil
		return
	}
	if node == this.tail {
		this.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
