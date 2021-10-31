package bfs

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
102. 二叉树的层序遍历
https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
 */

var q []*TreeNode
var ans [][]int

func levelOrder(root *TreeNode) [][]int {
	ans = [][]int{}
	q = []*TreeNode{}
	if root == nil {
		return ans
	}
	check(root)
	return ans
}

func check(root *TreeNode) {
	q = append(q, root)

	lay := 0
	for len(q) > 0 {
		n := len(q)
		for i := 0;i < n;i++ {
			h := q[0]
			q = q[1:]
			if len(ans) > lay {
				ans[lay] = append(ans[lay], h.Val)
			} else {
				ans = append(ans, []int{h.Val})
			}

			if h.Left != nil {
				q = append(q, h.Left)
			}

			if h.Right != nil {
				q = append(q, h.Right)
			}
		}
		//  遍历完一次说明第几层结束了
		lay++
	}
}