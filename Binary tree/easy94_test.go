package Binary_tree

/*
二叉树中序遍历
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	s := make([]int, 0)
	dfs := func(node *TreeNode) {}
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		s = append(s, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return s
}
