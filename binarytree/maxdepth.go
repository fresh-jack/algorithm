package binarytree

func maxDepthR(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxDepth(root *TreeNode) int {
	var (
		stack []*TreeNode
		result int
	)
	stack = append(stack, root)
	for len(stack) > 0 {
		var temp []*TreeNode
		for i := 0; i < len(stack); i++ {
			if stack[i].Left != nil {
				temp = append(temp, stack[i].Left)
			}
			if stack[i].Right != nil {
				temp= append(temp, stack[i].Right)
			}
		}
		stack = temp
		result++
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}