package main

import (
	"fmt"
	"github.com/thearyanahmed/libsandbox/array"
)

func main() {
	input := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}

	fmt.Println(array.InsertionSort(input))
}
