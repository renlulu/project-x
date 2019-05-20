package project_x

func merge2(nums1 []int, m int, nums2 []int, n int)  {
	left := m - 1
	right := n - 1
	merge := m + n -1

	for left >= 0 && right >= 0 {
		if nums1[left] > nums2[right] {
			nums1[merge] = nums1[left]
			left--
		} else {
			nums1[merge] = nums2[right]
			right--
		}
		merge--
	}

	for left >= 0 {
		nums1[merge] = nums1[left]
		left--
		merge--
	}

	for right >= 0 {
		nums1[merge] = nums2[right]
		right--
		merge--
	}
}
