/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    balanced := true

    var height func(*TreeNode) int
    height = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        left := height(node.Left)
        right := height(node.Right)

        if abs(left-right) > 1 {
            balanced = false
        }

        return 1 + max(left, right)
    }

    height(root)

    return balanced
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
