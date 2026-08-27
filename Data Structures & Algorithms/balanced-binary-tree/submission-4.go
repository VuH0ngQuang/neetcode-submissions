/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    return checkDepth(root) != -1
}

func checkDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := checkDepth(root.Left)
    if left == -1 {
        return -1
    }
    right := checkDepth(root.Right)
    if right == -1 {
        return -1
    }
    if abs(left-right) > 1 {
        return -1
    }
    return max(left, right) + 1
}

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}