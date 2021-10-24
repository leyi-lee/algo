package week03

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var inorderMap map[int]int // key val  val:index

func BuildTree(inorder []int, postorder []int) *TreeNode {
	inorderMap = map[int]int{}
	for key,val := range inorder {
		inorderMap[val] = key
	}
	return build(0, len(inorder) - 1, 0, len(postorder) - 1, inorder, postorder)
}

func build(l1 int, r1 int, l2 int, r2 int, inorder []int, postorder []int) *TreeNode {
	if l1 > r1 {
		return nil
	}
	root := &TreeNode{
		Val: postorder[r2],
	}

	mid,_ := inorderMap[root.Val]

	// l1 ~ mid - 1 左子树中序
	// mid +1 ~ r1 右子树中序

	/**
		中序遍历 inorder = [9,3,15,20,7]
		后序遍历 postorder = [9,15,7,20,3]
		因为后序遍历是左右根
		所以从右子树开始找，因为根再最后
		找出根后，那中序遍历 右边就是右子树 所以 为 mid + 1 -> r1
		那后序遍历的数组中如何找右子树开始的节点？
			找到个数 从中序遍历找 各位为 r1 - (mid + 1) + 1 = r1 - mid
		那后序遍历数组中的开始位置是 r2 - m  = r1 - mid 求m 为 r2 - r1 + mid

		那中序遍历的左子树开始为 l1 -> mid - 1
		那左子树后续遍历数组位置就是 l2 -> r2 - r1 + mid -1
		右子树的遍历
	 */

	root.Right = build(mid + 1, r1, r2 - r1 + mid, r2-1, inorder, postorder)
	root.Left = build(l1, mid - 1, l2, r2 - r1 + mid -1, inorder, postorder)
	return root
}
