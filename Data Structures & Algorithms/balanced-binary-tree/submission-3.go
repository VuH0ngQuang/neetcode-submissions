/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    if abs(depth(root.Left) - depth(root.Right)) > 1 {
        return false
    }
    return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(depth(root.Left), depth(root.Right)) + 1
}

func abs (num int) int {
    if num < 0 {
        return -num
    }
    return num
}
