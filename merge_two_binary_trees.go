package project_x

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	root := &TreeNode {
		Val: t1.Val + t2.Val,
		Left: mergeTrees(t1.Left,t2.Left),
		Right: mergeTrees(t1.Right,t2.Right),
	}

	return root
}
