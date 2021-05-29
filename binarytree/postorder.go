package binarytree

func postOrderR(root *TreeNode) []int {
	var (
		post func(*TreeNode)
		result []int
	)
	post = func(root *TreeNode) {
		if root == nil {
			return
		}
		post(root.Left)
		post(root.Right)
		result = append(result, root.Val)
	}
	post(root)
	return result
}

//
func postOrder(root *TreeNode) []int {
	var (
		stack []*TreeNode
		result []int
		pre *TreeNode
	)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if root.Right == nil || root.Right == pre {
			result = append(result, root.Val)
			pre = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return result
}
