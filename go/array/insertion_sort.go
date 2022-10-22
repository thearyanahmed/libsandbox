package array

func InsertionSort(arr []int) []int {
	l := len(arr)

	for i := 0; i < l; i++ {
		for j := i; j > 0 && arr[j -1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	return arr
}