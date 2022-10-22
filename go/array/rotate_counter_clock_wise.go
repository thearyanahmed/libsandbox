package array

// rotate by 2
// 1,2,3,4,5,6,7

// -> 3,4,5,6,7,1,2 # left rotate ?
//	 -> 3, 4, 5, 6, 7 -> 1, 2

func RotateCounterClockwise(arr []int32, n int32) []int32 {
	length := int32(len(arr))
	firstBlock := make([]int32,0)
	result := make([]int32, length - n)

	i := int32(0)

	for ; i < length ; i++ {
		if i < n {
			firstBlock = append(firstBlock, arr[i])
		} else {
			result[i - n] = arr[i]
		}
	}

	for _, v := range firstBlock {
		result = append(result, v)
	}

	return result
}