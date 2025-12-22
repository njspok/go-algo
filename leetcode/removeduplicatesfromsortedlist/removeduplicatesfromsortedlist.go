package removeduplicatesfromsortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// Complexity
// - O(n) by time
// - O(1) by space
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	for {
		if cur.Next == nil {
			return head
		}

		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
}
