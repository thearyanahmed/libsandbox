package main

import (
	"fmt"

	"github.com/thearyanahmed/rknot/linked_list"
)

func main() {

	lRUCache := linked_list.NewLRU(2)

	lRUCache.Put(1, 1) // cache is {1=1}
	lRUCache.Put(2, 2) // cache is {1=1, 2=2}
	// lRUCache.Get(1)    // return 1
	fmt.Printf("get(%v) = %v\n", 1, lRUCache.Get(1))
	lRUCache.PrintLURKey()

	lRUCache.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	// lRUCache.Get(2)    // returns -1 (not found)
	// fmt.Printf("get(%v) = %v\n", 2, lRUCache.Get(2))
	lRUCache.PrintLURKey()

	// lRUCache.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	// fmt.Printf("get(%v) = %v\n", 1, lRUCache.Get(1))
	// fmt.Printf("get(%v) = %v\n", 3, lRUCache.Get(3))
	// fmt.Printf("get(%v) = %v\n", 4, lRUCache.Get(4))

	// lRUCache.Get(1)    // return -1 (not found)
	// lRUCache.Get(3)    // return 3
	// lRUCache.Get(4)    // return 4

	// lRUCache.Put(1, 1)
	// lRUCache.Put(2, 2)
	// lRUCache.Put(3, 5)
	// lRUCache.Put(3, 6)
	lRUCache.PrintLRU()
	lRUCache.PrintDatastore()

	// key := 3
	// fmt.Printf("get(%v) = %v\n", key, lRUCache.Get(key))

}
