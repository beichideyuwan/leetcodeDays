package Binary_tree

/*
对称二叉树
给你一个二叉树的根节点 root ， 检查它是否轴对称。
*/
func isSymmetric(root *TreeNode) bool {
	return isEq(root, root)
}
func isEq(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isEq(left.Left, right.Right) && isEq(left.Right, right.Left)
}
