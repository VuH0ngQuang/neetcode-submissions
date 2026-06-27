/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func isSame (q *TreeNode, p *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}

	if q == nil || p == nil {
		return false
	}

	if q.Val != p.Val {
		return false
	}

	return isSame(q.Left, p.Left) && isSame (q.Right, p.Right)
 }

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if isSame(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
