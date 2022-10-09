package main

import (
	"fmt"

	"github.com/thearyanahmed/libsandbox/binary_search"
)

func main() {
	vec := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}

	m := make(map[int]bool)
	m[1] = true
	m[10] = true
	m[23] = true
	m[11] = true
	m[34] = true
	m[-12] = false
	m[123] = false
	// m[1111] = false

	for k, v := range m {
		got := binary_search.SearchIn2DMatrix(vec, k)

		fmt.Printf("%v == %v\n", got, v)
	}
}
