package array

import "sort"
func RemoveDuplicates(arr []int) []int {
	sort.Ints(arr)
	
	slow := 0

	for fast := 0; fast < len(arr) - 1; fast++ {
		if arr[slow] != arr[fast] {
			slow++
			arr[slow] = arr[fast]
		}
	}

	return arr[0:slow + 1]
}