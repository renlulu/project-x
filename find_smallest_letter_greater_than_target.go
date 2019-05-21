package project_x

func nextGreatestLetter(letters []byte, target byte) byte {
	left := 0
	right := len(letters) - 1

	for left <= right {
		middle := (left + right) / 2
		if letters[middle] <= target {
			left = middle + 1

		} else {
			right = middle - 1
		}
	}

	if left >= len(letters) {
		return letters[0]
	}

	return letters[left]
}
