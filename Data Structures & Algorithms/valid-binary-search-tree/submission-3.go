/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func valid (root *TreeNode, low *int, high *int) bool {
	if root == nil {
		return true
	}

	if low != nil && root.Val <= *low {
		return false
	}
	if high != nil && root.Val >= *high {
		return false
	}

	return valid(root.Left, low, &root.Val) && valid(root.Right, &root.Val, high)
}

func isValidBST(root *TreeNode) bool {
    return valid(root, nil, nil)
}
