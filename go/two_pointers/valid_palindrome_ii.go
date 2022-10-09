package two_pointers

func validPalindrome(s string) bool {
    valid, l, r := isPalindrome(s, 0, len(s)-1)
    if valid {
        return true
    }
    if valid, _, _ := isPalindrome(s, l+1, r); valid {
        return true
    }
    if valid, _, _ := isPalindrome(s, l, r-1); valid {
        return true
    }
    return false
}

func isPalindrome(s string, l, r int) (bool, int, int) {
    for l < r {
        if s[l] != s[r] {
            return false, l, r
        }
        l++
        r--
    }
    return true, 0, 0
}