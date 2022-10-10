package two_pointers

import "fmt"

func Trap(height []int) int {
	length := len(height)

	if length == 0 {
		fmt.Println("returning 0")
		return 0
	}

	l, r := 0, length-1
	lMax, rMax := height[l], height[r]
	res := 0

	for l < r {
		if lMax < rMax {
			l++
			lMax := getMax(lMax, height[l])

			res += lMax - height[l]
		} else {
			r--
			rMax := getMax(rMax, height[r])

			res += rMax - height[r]
		}

	}

	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
