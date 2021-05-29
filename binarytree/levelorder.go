package binarytree

// levelOrder ...
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		stack []*TreeNode
		result [][]int
	)
	stack = append(stack, root)
	for len(stack) > 0 {
		var (
			r []int
			temp []*TreeNode
		)
		for i := 0; i < len(stack); i++ {
			r = append(r, stack[i].Val)
			if root.Left != nil {
				temp = append(temp, stack[i].Left)
			}
			if root.Right != nil {
				temp = append(temp, stack[i].Right)
			}
		}
		stack = temp
		result = append(result, r)
	}
	return result
}
