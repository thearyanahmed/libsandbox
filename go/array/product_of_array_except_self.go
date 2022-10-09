package array

func ProductExceptSelf(arr []int) []int {
	length := len(arr)

	res := make([]int, length)

	prefix := 1

	for i := 0; i < length; i++ {
		res[i] = prefix
		prefix *= arr[i]
	}

	postfix := 1

	i := length - 1

	for ; i >= 0; i-- {
		res[i] *= postfix
		postfix *= arr[i]
	}

	return res
}
