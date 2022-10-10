package main

import (
	"fmt"

	"github.com/thearyanahmed/libsandbox/two_pointers"
)

func main() {
	fmt.Println(two_pointers.Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
