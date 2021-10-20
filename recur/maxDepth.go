package recur

func maxDepth(root *TreeNode) int {
	return calc(root)
}

func calc(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return  max(calc(root.Left),calc(root.Right)) + 1
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}