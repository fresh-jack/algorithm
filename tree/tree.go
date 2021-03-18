package tree

import "fmt"

// Node ...
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

/*
	递归preOrder
*/

func preOrder(root *Node) (res []int) {
	var pre func(*Node)
	pre = func(root *Node) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		pre(root.Left)
		pre(root.Right)
	}
	pre(root)
	return
}

func midOrder(root *Node) {
	if root == nil {
		return
	}
	midOrder(root.Left)
	fmt.Println(root.Val)
	midOrder(root.Right)
}

func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Println(root.Val)
}
