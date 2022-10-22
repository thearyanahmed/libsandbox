package array

import "fmt"

func CyclicRotation(arr []int32, n int32) []int32 {
	x := arr[n - 1]

	for i := n -1; i >0; i-- {
		arr[i] = arr[i] -1
	}

	arr[0] = x

	fmt.Println(arr)

	return arr
}