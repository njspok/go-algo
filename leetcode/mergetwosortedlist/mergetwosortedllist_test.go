package mergetwosortedlist

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// todo implement
	return nil
}

func Test(t *testing.T) {
	tests := []struct {
		a      *ListNode
		b      *ListNode
		result []int
	}{
		{a: nil, b: nil, result: []int{}},
		{a: nil, b: nil, result: []int{}},
	}
	for n, tt := range tests {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			res := mergeTwoLists(tt.a, tt.b)
			actual := toSlice(res)
			require.Equal(t, tt.result, actual)
		})
	}
}

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

func toSlice(head *ListNode) []int {
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

func TestMakeList(t *testing.T) {
	tests := [][]int{
		{},
		{1},
		{1, 2, 3, 4, 5},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			list := makeList(test)
			actual := toSlice(list)
			require.Equal(t, test, actual)
		})
	}
}
