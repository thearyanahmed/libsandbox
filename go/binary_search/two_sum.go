package binary_search

func TwoSum(numbers []int, target int) []int {

	r := len(numbers) - 1

	for i := range numbers {

		l := i + 1
		t := target - numbers[i]

		for l <= r {

			mid := l + (r-l)/2

			if numbers[mid] == t {
				res := []int{i + 1, mid + 1}

				return res
			}

			if numbers[mid] < t {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	// x := []int{}
	return []int{}
}
