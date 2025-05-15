package mergetwosortedlist

func makeList(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}

	res := make([]*ListNode, 0, len(s))

	for i := 0; i < len(s); i++ {
		if i == 0 {
			res = append(res, &ListNode{
				Val:  s[i],
				Next: nil,
			})
			continue
		}

		res = append(res, &ListNode{
			Val:  s[i],
			Next: nil,
		})

		res[i-1].Next = res[i]
	}

	return res[0]
}

func copyList(list *ListNode) *ListNode {
	var res *ListNode
	prev := res

	cur := list
	for cur != nil {
		newNode := &ListNode{
			Val:  cur.Val,
			Next: nil,
		}

		if prev == nil {
			res = newNode
			prev = res
		} else {
			prev.Next = newNode
			prev = prev.Next
		}

		cur = cur.Next
	}

	return res
}

func listToSlice(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
