package project_x

import "strconv"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	return isPalindrome2(strconv.FormatInt(int64(x),10))

}


func isPalindrome2(s string) bool {

	if len(s)%2 == 0 {
		l := len(s)
		half := l / 2

		for i := 0; i < half; i++ {
			if (s[i] != s[l-1-i]) {
				return false
			}
		}

	}

	if len(s)%2 != 0 {
		l := len(s)
		half := l / 2

		for i := 0; i < half+1; i++ {
			if (s[i] != s[l-1-i]) {
				return false
			}
		}
	}

	return true
}
