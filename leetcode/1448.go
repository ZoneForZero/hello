package lc

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	return goodNode(root, -10005)
}
func goodNode(root *TreeNode, maxVal int) int {
	if root == nil {
		return 0
	}
	ans := 0
	if root.Val >= maxVal {
		ans++
		maxVal = root.Val
	}
	return ans + goodNode(root.Left, maxVal) + goodNode(root.Right, maxVal)
}
