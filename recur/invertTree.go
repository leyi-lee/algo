package recur

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
226. 翻转二叉树
https://leetcode-cn.com/problems/invert-binary-tree/description/
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	InvertTree(root.Left)
	InvertTree(root.Right)
	return root
}