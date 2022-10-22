package array

import "fmt"

// Problem:
// arr[] = {1, 2, 3, 4, 5, 6, 7}, n = 2
// Output: 3 4 5 6 7 1 2

// Solution 1:
// take from n to last, create a new array
// append the first to nth values
func RotateArrayByNPosition(arr []int32, n int) []int32 {
	length := len(arr)
	res := make([]int32,0,length)

	for i := n; i < length; i++ {
		res = append(res, arr[i])
	}

	for i := 0; i < n; i++ {
		res = append(res,arr[i])
	}

	return res
}
