package binarytree

// 递归
func preOrderR(root *TreeNode) []int {
	var (
		pre func(*TreeNode)
		result []int
	)

	pre = func(root *TreeNode) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		pre(root.Left)
		pre(root.Right)
	}
	pre(root)
	return  result
}

// 非递归
func preOrder(root *TreeNode) []int {
	var (
		result []int
		stack []*TreeNode
	)
	for root != nil || len(stack) > 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return result
}
