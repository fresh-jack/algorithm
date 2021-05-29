package binarytree

func inOrderR(root *TreeNode) []int {
	var (
		in  func(node *TreeNode)
		result []int
	)
	in = func(root *TreeNode) {
		if root == nil {
			return
		}
		in(root.Left)
		result = append(result, root.Val)
		in(root.Right)
	}
	in(root)
	return result
}

func inOrder(root *TreeNode) []int {
	var (
		stack []*TreeNode
		result []int
	)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		root = root.Right
	}

	return result
}