package main

import (
	"fmt"
	"github.com/thearyanahmed/libsandbox/array"
	"github.com/thearyanahmed/libsandbox/utils"
)

func main() {
	input := utils.RandomSlice(20)

	fmt.Println("Input ->",input)

	fmt.Println(array.InsertionSort(input))
}
