package two_pointers

func MaxArea(height []int) int {
	res := 0

	left, right := 0, len(height)-1

	for left < right {
		area := (right - left) * min(height[left], height[right])
		res = max(res, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
