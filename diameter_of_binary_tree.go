package project_x

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

var m int

func diameterOfBinaryTree(root *TreeNode) int {
	m = 0
	maxDepth2(root)
	return m
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth2(root.Left)
	right := maxDepth2(root.Right)

	m = max(m, left+right)

	return max(left, right) + 1
}
