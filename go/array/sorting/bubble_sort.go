package array

func BubbleSort(arr []int) []int{
	swapped := true

	l := len(arr)

	for swapped { // run
		swapped = false // pause, if i >= l, breaks

		for i := 1; i < l ; i++ { // from 1th to last, keep checking if arr[i -1] > 1 or not
			if arr[i - 1] > arr[i] {
				arr[i], arr[i - 1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}

	return arr
}