package project_x

import "math"

func judgeSquareSum(c int) bool {
	i := 0
	j := int(math.Sqrt(float64(c)))

	for i <= j {
		r := i*i + j*j
		if c == r {
			return true
		}

		if r > c {
			j--
		}

		if r < c {
			i++
		}
	}

	return false

}
