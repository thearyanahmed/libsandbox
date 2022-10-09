package binary_search

func SearchIn2DMatrix(matrix [][]int, target int) bool {

	lengthOfMatrix := len(matrix)

	if lengthOfMatrix == 0 {
		return false
	}

	lengthOfRow := len(matrix[0])

	// loop through each row
	// loop through each column
	// find the upper and lower found of that row

	left, right := 0, lengthOfMatrix-1

	for left <= right {
		if target > matrix[left][lengthOfRow-1] {
			left++
			continue
		}

		for i := 0; i < lengthOfRow; i++ {
			if matrix[left][i] == target {
				return true
			}
		}

		left++
	}

	return false
}
