package main

import (
	"fmt"

	"github.com/thearyanahmed/libsandbox/array"
)

func main() {
	vec := []int{1, 2, 3, 4, 100, 200}

	fmt.Println(array.LongestConsecutiveSequnce(vec))
}
