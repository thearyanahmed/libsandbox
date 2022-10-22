package array

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

// rotate by 2
// 1,2,3,4,5,6,7

// -> 7,6,1,2,3,4,5 # right rotate ?
// 	-> Take the last 2 elements
//  -> Reverse it 7, 6, -> 1,2,3,4,5 = 7, 6, 1, 2, 3, 4, 5

// -> 3,4,5,6,7,1,2 # left rotate ?
//	 -> 3, 4, 5, 6, 7 -> 1, 2


// reverse
// in place swap 