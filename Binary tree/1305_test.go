package Binary_tree

/*
给你 root1 和 root2 这两棵二叉搜索树。请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.
*/
func inorder(root *TreeNode) []int {
	var res []int
	dfs := func(root *TreeNode) {}
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return res
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	r1 := inorder(root1)
	r2 := inorder(root2)
	head1, end1 := 0, len(r1)
	head2, end2 := 0, len(r2)
	merged := make([]int, 0, end1+end2)
	for {
		if head1 == end1 {
			return append(merged, r2[head2:]...)
		}
		if head2 == end2 {
			return append(merged, r1[head1:]...)
		}
		if r1[head1] < r2[head2] {
			merged = append(merged, r1[head1])
			head1++
		} else {
			merged = append(merged, r2[head2])
			head2++
		}
	}
}
