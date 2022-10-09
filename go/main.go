package main

import (
	"fmt"

	"github.com/thearyanahmed/libsandbox/binary_search"
)

func main() {
	vec := []int{3, 6, 7, 11}

	fmt.Println(binary_search.KokoEatingBananas(vec, 8))
}
