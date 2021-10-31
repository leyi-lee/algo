package week04

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
538. 把二叉搜索树转换为累加树
https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
 */

/**
先遍历遍历右子树，然后累加值
 */

var sum int
func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	traverse(root.Right)
	sum += root.Val
	root.Val = sum
	traverse(root.Left)
}
