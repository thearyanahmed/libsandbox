package array

// SelectionSort
// 1. for each element
// 2. minIndex = i
// 3. for each element from the next one (j = i + 1)
// 4. if current element is < minIndex , swap
func SelectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}

	return arr
}