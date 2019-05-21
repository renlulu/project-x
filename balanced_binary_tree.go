package project_x

import "strings"

var result bool

func isBalanced(root *TreeNode) bool {
	result = true
	maxDepth1(root)
	strings.LastIndexByte("ddd",'d')
	return result
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	if (l-r) > 1 || (r-l) > 1 {
		result = false
	}

	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
