package project_x

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	left := 1
	right := x

	for left <= right {
		middle := (left + right) / 2

		if middle * middle == x {
			return middle
		}

		if middle * middle > x {
			right = middle - 1
			continue
		}

		if middle * middle < x {
			left = middle + 1
		}
	}

	return left-1
}