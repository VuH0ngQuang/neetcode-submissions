/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    return 1 + dfs(root.Left, root.Val) + dfs(root.Right, root.Val)
}

func dfs(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}
	if max <= root.Val {
		max = root.Val
		return 1 + dfs(root.Left, max) + dfs(root.Right, max)
	}
	return dfs(root.Left, max) + dfs(root.Right, max)

}