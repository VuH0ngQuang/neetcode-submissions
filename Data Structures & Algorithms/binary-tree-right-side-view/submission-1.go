/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type StackNode struct {
	node *TreeNode
	depth int
 }

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res = []int{}
	var stack = []StackNode{{root, 0}}
    
	for len(stack) > 0 {
		n := len(stack)
		item := stack[n-1]
		stack = stack[:n-1]
		
		node, depth := item.node, item.depth

		if depth == len(res) {
			res = append(res, node.Val)
		}
		if node.Left != nil {
			stack = append(stack, StackNode{node.Left, depth + 1})
		}
		if node.Right != nil {
			stack = append(stack, StackNode{node.Right, depth + 1})
		}
	}

	return res
}
