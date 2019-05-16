package project_x

func LongestCommonPrefix(strs []string) string {
	end := -1

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	first := strs[0]
	others := strs[1:]

	for i := 0; i < len(first); i++ {
		if allHaveOn(others, first[i], i) {
			end += 1
		} else {
			break
		}
	}

	if end == -1 {
		return ""
	}

	return first[0:end+1]
}

func allHaveOn(all []string, c byte, i int) bool {
	for j := 0; j < len(all); j++ {
		if !haveOn(all[j], c, i) {
			return false
		}
	}

	return true
}

func haveOn(s string, c byte, i int) bool {
	if len(s) <= i {
		return false
	}
	return s[i] == c
}
