package array

import "fmt"

func QuickSort(arr []int, low, high int) []int {
	fmt.Println("todo!()")
	return arr
	if low < high {
		pivot := parition(arr, low, high)

		QuickSort(arr, low, pivot)
		QuickSort(arr, pivot + 1, high)
	}

	return arr
}

func parition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low
	j := high

	for i < j {
		if arr[i] <= pivot && i < high {
			i++
		}

		if arr[j] > pivot && j > low {
			j++
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}


	arr[low] = j
	arr[j] = pivot

	return j
}