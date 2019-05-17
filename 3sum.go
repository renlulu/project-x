package project_x

import "sort"

//hack..
func all0(m []int) bool {
	if len(m) < 3 {
		return false
	}
	for i := 0; i < len(m); i++ {
		if m[i] != 0 {
			return false
		}
	}

	return true
}

func isExist(m [][]int, n []int) bool {
	for i := 0; i < len(m); i++ {
		k := m[i]
		if equal(k, n) {
			return true
		} else {
			continue
		}
	}

	return false
}

func equal(m []int, n []int) bool {
	if len(m) != len(n) {
		return false
	}

	for i := 0; i < len(m); i++ {
		if m[i] != n[i] {
			return false
		}
	}

	return true
}

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if all0(nums) {
		r := []int{0, 0, 0}
		result = append(result, r)
		return result
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			S := make(map[int]int, len(nums))
			for k := 0; k < len(nums); k++ {
				if k == i || k == j {
					continue
				}
				_, ok := S[nums[k]]
				if ok {
					S[nums[k]] += 1
				} else {
					S[nums[k]] = 1
				}
			}
			target := 0 - (nums[i] + nums[j])
			_, exist := S[target]
			if exist {
				r := []int{nums[i], nums[j], target}
				sort.Ints(r)
				if !isExist(result, r) {
					result = append(result, r)
				}

			}
		}

	}

	return result
}
