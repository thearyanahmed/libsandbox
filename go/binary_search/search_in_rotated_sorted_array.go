package binary_search

func SearchInRotatedSortedArray(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	if low > high {
		return -1
	}
	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		}

		// 15,16,17,18,19,1,2,3,4,5,6,7,8,9,10
		if arr[low] < arr[mid] {
			// mid to high is sorted
		} else if arr[mid] < target && target < arr[high] { // mid to high is sorted, target is in this range

		}


	}
	return -1
}