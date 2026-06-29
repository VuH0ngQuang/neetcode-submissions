/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs (n *TreeNode, count *int, m int) {
	if n == nil {
		return
	}
	if n.Val >= m {
		*count++
		m = max(m, n.Val)
	}
	if n.Left != nil {
		dfs(n.Left, count, m)
	}
	if n.Right != nil {
		dfs(n.Right, count, m)
	}
} 

func goodNodes(root *TreeNode) int {
    count := 0
	dfs(root, &count, root.Val)

	return count
}
