package main

import (
	"fmt"

	"github.com/thearyanahmed/libsandbox/array"
)

func main() {
	vec := []int{-1, 1, 0, -3, 3}

	fmt.Println(array.ProductExceptSelf(vec))
}
