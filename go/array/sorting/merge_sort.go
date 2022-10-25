package array

// MergeSort
// 1. if length of the array is less than 2, return arr
// 2. else, take 2 halves, 1:1
// 3. return merge(1st : 2nd)
func MergeSort(arr []int) []int {
	l := len(arr)

	if l < 2 {
		return arr
	}

	first := MergeSort(arr[:l/2])
	last := MergeSort(arr[l/2:])

	return merge(first, last)
}

// merge
// 1. merge the both for based on order
// 2. DON'T REASSIGN i or j
// 3. append and return res
func merge (first, last []int) []int {
	var res []int
	i, j := 0, 0

	lengthOfFirst := len(first)
	lengthOfLast := len(last)

	for i < lengthOfFirst && j < lengthOfLast {
		if first[i] < last[j] {
			res = append(res, first[i])
			i++
		} else {
			res = append(res, last[j])
			j++
		}
	}

	for ; i < lengthOfFirst; i++ {
		res = append(res, first[i])
	}

	for ; j < lengthOfLast; j++ {
		res = append(res, last[j])
	}

	return res
}