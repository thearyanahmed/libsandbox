package array

func Reverse(arr []int32) []int32 {
	length := len(arr)

	mid := length / 2

	i := 0
	last := length - 1

	for ; i < mid; {
		arr[i], arr[last] = arr[last], arr[i]

		i++
		last--
	}

	return arr
}