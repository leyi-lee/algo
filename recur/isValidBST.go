package recur

/**
98. 验证二叉搜索树
https://leetcode-cn.com/problems/validate-binary-search-tree/
*/

func IsValidBST(root *TreeNode) bool {
	return check(root, -10000000, 1000000)
}

func check(root *TreeNode, rangeLeft int, rangeRight int) bool {
	if root == nil {
		return true
	}

	if root.Val < rangeLeft || root.Val > rangeRight {
		return false
	}

	return check(root.Left, rangeLeft, root.Val - 1) && check(root.Right, root.Val + 1, rangeRight)
}
