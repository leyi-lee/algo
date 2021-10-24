package tree

/**
105. 从前序与中序遍历序列构造二叉树
https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
 */

/**
preorder 前序遍历
inorder 中序遍历
前序遍历 根左右，所以前序的第一个值就是根
找左子树  先找到中序的根，等于root.val就找到了，然后小于的就是左子树，大于的就是右子树
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(0, len(preorder) - 1, 0, len(inorder) - 1, preorder, inorder)
}

func build(l1 int, r1 int, l2 int, r2 int, preorder []int, inorder []int) *TreeNode {
	if l1 > r1 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[l1],
	}

	mid := l2 // 是在中序inorder中的位置
	for inorder[mid] != root.Val {
		mid++
	}

	// l2 ~ mid - 1 左子树中序
	// mid +1 ~ r2 右子树中序

	/**
	 解释
		前序preorder
		l1 + 1就是除去根，下一个 的开始
		那左子树下面有多少个节点，得靠中序去求，左子树的中序=l2 ~ mid - 1,所以个数为 mid-1-l2+1
		例如 [0,1,2,3,4] 4-2+1 = 3 下标和值举例一样
		l1 + m = l1 + (mid - 1) - l2 + 1
	 */
	root.Left = build(l1 + 1, l1 + (mid - 1) - l2 + 1, l2, mid -1, preorder, inorder)
	root.Right = build(l1 + (mid - 1) - l2 + 1 + 1, r1, mid + 1, r2, preorder, inorder)

	return root
}

