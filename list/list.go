package list

// Node ...
type Node struct {
	Val  int
	Next *Node
}

/*
	reverse List
*/
func reverseList(head *Node) *Node {
	var prev *Node
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = prev

		cur = temp
	}

	return prev
}

/*
	middleNode
*/
func middleNode(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

/*
	remove duplicate nodes
*/
func removeDuplicateNodes(head *Node) *Node {
	cur := head
	for cur != nil {
		temp := cur
		for temp.Next != nil {
			if temp.Next.Val == cur.Val {
				temp.Next = temp.Next.Next
			} else {
				temp = temp.Next
			}
		}
		cur = cur.Next
	}

	return head
}

func removeDuplicateNodes2(head *Node) *Node {
	if head != nil {
		cur := head
		m := map[int]struct{}{head.Val: struct{}{}}
		for cur.Next != nil {
			if _, ok := m[cur.Next.Val]; !ok {
				m[cur.Next.Val] = struct{}{}
				cur = cur.Next
			} else {
				cur.Next = cur.Next.Next
			}
		}
	}
	return head
}
