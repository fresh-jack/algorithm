package binarytree

func pathSum(root *TreeNode, targetNum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetNum == root.Val
	}
	return pathSum(root.Left, targetNum - root.Val) || pathSum(root.Right, targetNum - root.Val)
}

