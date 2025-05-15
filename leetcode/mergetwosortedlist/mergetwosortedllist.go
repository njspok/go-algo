package mergetwosortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		if list2 == nil {
			return nil
		}
		return list2
	} else if list2 == nil {
		return list1
	}

	res := copyList(list1)

	cur := list2

	for cur != nil {
		res = placeNode(cur, res)
		cur = cur.Next
	}

	return res
}

func placeNode(node *ListNode, list *ListNode) *ListNode {
	cur := list
	prev := (*ListNode)(nil)

	for cur != nil {
		// in begin
		if prev == nil {
			if node.Val <= cur.Val {
				newNode := &ListNode{
					Val:  node.Val,
					Next: cur,
				}
				list = newNode
				break
			}
		}

		if node.Val <= cur.Val {
			newNode := &ListNode{
				Val:  node.Val,
				Next: cur,
			}
			prev.Next = newNode
			break
		} else {
			if cur.Next == nil {
				newNode := &ListNode{
					Val:  node.Val,
					Next: nil,
				}
				cur.Next = newNode
				break
			}
		}

		prev = cur
		cur = cur.Next
	}

	return list
}
