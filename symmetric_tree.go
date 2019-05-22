package project_x

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricHelper(root.Left,root.Right)
}

func isSymmetricHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return isSymmetricHelper(left.Left,right.Right) || isSymmetricHelper(left.Right,right.Left)
}
