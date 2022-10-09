package array

func LongestConsecutiveSequnce(nums []int) int {
	m := map[int]bool{}

	for _, num := range nums {
		m[num] = true
	}

	longest := 0
	for num := range m {
		if m[num-1] {
			continue
		}

		cur := num
		for m[cur+1] {
			cur++
		}

		longest = max(longest, cur-num+1)
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
