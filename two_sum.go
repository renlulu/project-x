package project_x

func TwoSum(nums []int, target int) []int {
	for k := 0; k < len(nums); k++ {
		for p := k + 1; p < len(nums); p++ {
			if nums[k]+nums[p] == target {
				return []int{k, p}
			}
		}
	}
	return nil
}
