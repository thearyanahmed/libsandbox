package binary_search

func SquareRoot(n int32) int32 {
	if n == 0 {
		return 0
	}

	var left int32 = 0
	right := n

	var res int32 = -1

	for left <= right {
		mid := left + ((right - left) / 2)

		if mid*mid <= n {
			left = mid + 1
			res = mid
		} else {
			right = mid - 1
		}
	}

	return res
}
