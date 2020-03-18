/*
 * Package dimaoftree 二叉树的最大直径
 * @Author: wang shouli
 * @LastEditors: wang shouli
 */

package dimaoftree

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

func maxdepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftDepth = 0
	var rightDepth = 0
	if root.Left != nil {
		leftDepth = 1 + maxdepth(root.Left)
	}
	if root.Right != nil {
		rightDepth = 1 + maxdepth(root.Right)
	}

	if (leftDepth + rightDepth) > max {
		max = leftDepth + rightDepth
	}

	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ret = 0
	if root.Left != nil {
		ret += (1 + maxdepth(root.Left))
	}
	if root.Right != nil {
		ret += (1 + maxdepth(root.Right))
	}
	if max > ret {
		return max
	}
	return ret
}
