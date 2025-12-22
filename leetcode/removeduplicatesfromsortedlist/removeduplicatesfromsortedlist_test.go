package removeduplicatesfromsortedlist

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		list []int
		res  []int
	}{
		{list: nil, res: nil},
		{list: []int{}, res: nil},
		{list: []int{1}, res: []int{1}},
		{list: []int{1, 2}, res: []int{1, 2}},
		{list: []int{1, 2, 3}, res: []int{1, 2, 3}},
		{list: []int{1, 1, 1}, res: []int{1}},
		{list: []int{1, 2, 2}, res: []int{1, 2}},
		{list: []int{1, 1, 2, 2, 3, 3, 4}, res: []int{1, 2, 3, 4}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.list), func(t *testing.T) {
			linkedList := makeLinkedList(test.list)
			linkedList = deleteDuplicates(linkedList)
			require.Equal(t, test.res, toSlice(linkedList))
		})
	}
}

func makeLinkedList(list []int) *ListNode {
	var head, curr *ListNode
	for _, v := range list {
		newNode := &ListNode{Val: v}
		if curr != nil {
			curr.Next = newNode
		} else {
			head = newNode
		}
		curr = newNode
	}
	return head
}

func toSlice(list *ListNode) []int {
	if list == nil {
		return nil
	}

	var res []int
	next := list
	for next != nil {
		res = append(res, next.Val)
		next = next.Next
	}

	return res
}
