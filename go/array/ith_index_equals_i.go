package array

import "fmt"

func IthIndexEqualsI(arr []int) []int {
	m := make(map[int]int)
	length := len(arr)

	for i := 0; i < length; i++ {
		if arr[i] == - 1 {
			continue
		}
		m[arr[i]] = i
	}

	for i := 0; i < length; i++ {
		if _, ok := m[i]; ok {
			arr[i] = i
		} else {
			arr[i] = -1
		}
	}

	return arr
}

func IthIndexEqualsWithoutHashmap(arr []int) []int {
	for i := 0; i < len(arr); {
		if arr[i] >= 0 && arr[i] != i { // arr[5] != 5
			arr[arr[i]], arr[i] = arr[i], arr[arr[i]]
		} else {
			i++
		}
	}

	return arr
}