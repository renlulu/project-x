package project_x

import "strings"

func reverseVowels(s string) string {
	vowels := map[string]interface{}{
		"a": nil,
		"e": nil,
		"i": nil,
		"o": nil,
		"u": nil,
		"A": nil,
		"E": nil,
		"I": nil,
		"O": nil,
		"U": nil,
	}

	result := make([]byte, len(s))

	first := 0
	last := len(s) - 1

	for first <= last {
		_, ok1 := vowels[string(s[first])]
		if !ok1 {
			result[first] = s[first]
			first++
			continue
		}

		_, ok2 := vowels[string(s[last])]
		if !ok2 {
			result[last] = s[last]
			last--
			continue
		}

		if ok1 && ok2 {
			result[first], result[last] = s[last], s[first]
			first++
			last--
			continue
		}
	}

	var sb strings.Builder
	sb.Write(result)
	return sb.String()
}
