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

func Test(t *testing.T) {
	tests := []struct {
		a      []int
		b      []int
		result []int
	}{
		{a: nil, b: nil, result: []int{}},
		{a: []int{1, 2, 3}, b: nil, result: []int{1, 2, 3}},
		{a: nil, b: []int{4, 5, 6}, result: []int{4, 5, 6}},
		{a: []int{1, 2, 3}, b: []int{0}, result: []int{0, 1, 2, 3}},
		{a: []int{1, 2, 3}, b: []int{1}, result: []int{1, 1, 2, 3}},
		{a: []int{1, 2, 3}, b: []int{2}, result: []int{1, 2, 2, 3}},
		{a: []int{1, 2, 3}, b: []int{3}, result: []int{1, 2, 3, 3}},
		{a: []int{1, 2, 3}, b: []int{4}, result: []int{1, 2, 3, 4}},
		{a: []int{1, 2, 3}, b: []int{4, 5}, result: []int{1, 2, 3, 4, 5}},
		{a: []int{1, 2, 3}, b: []int{4, 5, 6}, result: []int{1, 2, 3, 4, 5, 6}},
		{a: []int{4, 5, 6}, b: []int{1, 2, 3}, result: []int{1, 2, 3, 4, 5, 6}},
		{a: []int{1, 2, 4}, b: []int{1, 3, 4}, result: []int{1, 1, 2, 3, 4, 4}},
	}
	for n, tt := range tests {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			res := mergeTwoLists(makeList(tt.a), makeList(tt.b))
			actual := listToSlice(res)
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

func Test_makeList(t *testing.T) {
	tests := [][]int{
		{},
		{1},
		{1, 2, 3, 4, 5},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			list := makeList(test)
			actual := listToSlice(list)
			require.Equal(t, test, actual)
			require.Len(t, actual, len(test))
		})
	}
}

func Test_copyList(t *testing.T) {
	tests := []struct {
		list   []int
		result []int
	}{
		{list: nil, result: []int{}},
		{list: []int{1}, result: []int{1}},
		{list: []int{1, 2, 3}, result: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.list), func(t *testing.T) {
			list := makeList(tt.list)
			res := copyList(list)
			actual := listToSlice(res)
			require.Equal(t, tt.result, actual)
			require.Len(t, actual, len(tt.list))
		})
	}
}

func Test2(t *testing.T) {
	a := (*ListNode)(nil)
	fmt.Println(a == nil)
}
