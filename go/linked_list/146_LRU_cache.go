package linked_list

import "fmt"

type Cache struct {
	key      int
	value    int
	previous *Cache
	next     *Cache
}

type LRUCache struct {
	datastore         map[int]int
	capacity          int
	currently_holding int
	history           *Cache
}

func NewLRU(capacity int) LRUCache {
	ds := make(map[int]int)

	return LRUCache{
		datastore:         ds,
		capacity:          capacity,
		currently_holding: 0,
		history:           nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.datastore[key]; ok {
		// this.SwapLastHistory()
		return val
	}

	c := this.history.next

	c = c.next
	this.history = c

	// this.
	return -1
}

func (this *LRUCache) SwapLastHistory() {
	// 1 -> 2 -> 3

	nextKey, nextVal := this.history.next.key, this.history.next.value

	this.history.next.key = this.history.key
	this.history.next.value = this.history.value

	this.history.key = nextKey
	this.history.value = nextVal
}

// working
func (this *LRUCache) EvictLastHistory() {
	// 1-> 2-> 3->4
	c := this.history
	// x -> 2 -> 3 -> 4
	c = c.next // one step foreward
	c.previous = nil

	this.history = c
}

func (this *LRUCache) UpdateLastHistory(key, value int) {
	this.history.key = key
	this.history.value = value
}

func (this *LRUCache) UpdateCacheWhere(key, value int) {
	cache := this.history

	for cache != nil { // traverse the whole history to find the value
		if cache.key == key {
			// found the cache
			// swap it with head
			// swapping won't work: [we are swaping the head by swaping the keys and values]

			// cache.key = this.history.key
			// cache.value = this.history.value

			// take the previous key
			// take the next key
			prev := cache.previous
			next := cache.next

			prev.next = next
			next.previous = prev

			break
		}

		cache = cache.next
	}

	// finally add a new head head
	c := Cache{
		key:      key,
		value:    value,
		previous: nil,
		next:     this.history,
	}

	this.history = &c
}

func (this *LRUCache) Put(key int, value int) {
	// firstly, check if it exists already
	// if it does, we need to update the key and value
	if _, ok := this.datastore[key]; ok {

		// update of key-value is being done at the very end
		// need to update the history nodes

		// secondly, we need to take the node, bring it all way front,
		// sew the remaining parts

		// 1 -> 2 -> 3 -> 4
		// 3 -> 1 -> 2 -> 4

		this.UpdateCacheWhere(key, value)

	} else {
		// check if current holding + 1 > capacity
		// 		if yes,
		//			1. delete the last node, // we dont need to delete the node
		// 			2. create a new node
		// 			3. set this new node as head
		// 			4. attach the rest of history
		//			5. this.datastore[key] = value (maybe defer this?)
		//			6. NOPE, because we are deleting 1 node, so current_holding++ and -- [this.currently_holding++]

		if this.currently_holding+1 > this.capacity {
			// this.UpdateLastHistory(key, value)
			fmt.Printf("%d should be deleted from map\n", this.history.key)
			delete(this.datastore, this.history.key)

			this.EvictLastHistory()
			this.addToCache(key, value, this.history)
			// delete from map
		} else {
			// 		if no, 2- > 3 -> 4 -> 5 -> 6
			this.addToCache(key, value, this.history)
			this.currently_holding++
		}
	}

	this.datastore[key] = value
}

func (this *LRUCache) addToCache(key, value int, next *Cache) {
	//  1 -> 2 -> 3
	h := this.history

	for h != nil {
		h = h.next
	}

	fmt.Println("h", this.history)
}

func (this *LRUCache) PrintLRU() {
	fmt.Println("printing lru history")

	node := this.history

	for node != nil {
		fmt.Println(node.key, node.value)
		node = node.next
	}
}

func (this *LRUCache) PrintDatastore() {
	fmt.Println("printing datastore")

	for k, v := range this.datastore {
		fmt.Printf("%d = %d\n", k, v)
	}
}

func (this *LRUCache) PrintLURKey() {
	fmt.Println("lru key", this.history.key)
}
