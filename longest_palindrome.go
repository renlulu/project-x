package project_x

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	longest := 0
	start := 0
	end := 0

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if (isPalindrome(s[i : j+1])) {
				if (j+1-i > longest) {
					longest = j + 1 - i
					start = i
					end = j + 1
				}
			}

		}
	}

	result := s[start:end]
	if result == "" {
		return s[0:1]
	}

	return s[start:end]
}

func isPalindrome(s string) bool {

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
