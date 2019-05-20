package project_x

func validPalindrome(s string) bool {
	first := 0
	last := len(s) - 1

	for first < last {
		if s[first] != s[last] {
			return isPalindrome3(s, first+1, last) || isPalindrome3(s, first, last-1)
		}
		first++
		last--
	}

	return true
}

func isPalindrome3(s string, first, last int) bool {
	for first < last {
		if s[first] != s[last] {
			return false
		}
		first++
		last--
	}

	return true
}
