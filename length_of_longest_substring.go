package project_x

func LengthOfLongestSubstring(s string) int {
	longest := 0

	for i:=0; i<len(s);i++{
		characters := make(map[byte]int,10)
		count := 0

		for j:=i;j<len(s);j++{
			_,ok := characters[s[j]]
			if !ok {
				characters[s[j]] = 1
				count += 1
			} else {
				break
			}
		}

		if count > longest {
			longest = count
		}
	}

	return longest
}
