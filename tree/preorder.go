package tree

/**
589. N 叉树的前序遍历
https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
 */

type Node struct {
	Val int
	Children []*Node
}

var res []int

func preorder(root *Node) []int {
	res = []int{}
	pRecur(root)
	return res
}

func pRecur(root *Node)  {
	if root == nil {
		return
	}

	res = append(res, root.Val)
	for _, child := range root.Children {
		pRecur(child)
	}
}