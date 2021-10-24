package tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
94. 二叉树的中序遍历
https://leetcode-cn.com/problems/binary-tree-inorder-traversal/submissions/
 */

var ans []int
func inorderTraversal(root *TreeNode) []int {
	ans = []int{}
	recur(root)
	return ans
}

func recur(root *TreeNode)  {
	if root == nil {
		return
	}

	recur(root.Left)
	ans = append(ans, root.Val)
	recur(root.Right)
}