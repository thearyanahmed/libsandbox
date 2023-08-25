package two_pointers

func validPalindromeRemovingOneElement(s string) bool {
	isPalindrome := func(left, right int) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			// Try removing either s[left] or s[right]
			return isPalindrome(left+1, right) || isPalindrome(left, right-1)
		}
		left++
		right--
	}
	return true
}
