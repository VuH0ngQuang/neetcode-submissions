/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	count := 0
	return inorder(root, k, &count)
}

func inorder(root *TreeNode, k int, count *int) int {
	if root == nil {
		return -1
	}
	if res := inorder(root.Left, k, count); res != -1 {
		return res
	}
	*count++
	if *count == k {
		return root.Val
	}
	return inorder(root.Right, k, count)
}