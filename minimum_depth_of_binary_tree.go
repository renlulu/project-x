package project_x

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if l == 0 || r == 0 {
		return l + r + 1
	}

	if l < r {
		return l + 1
	} else {
		return r + 1
	}
}
