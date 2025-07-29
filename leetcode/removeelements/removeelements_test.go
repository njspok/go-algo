package removeelements

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"slices"
	"sort"
	"testing"
)

func removeElementTrivial(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] != val {
			i++
			continue
		}

		nums = append(nums[:i], nums[i+1:]...)
	}

	return len(nums)
}

func removeElement(nums []int, val int) int {

	dropOneElement := func(list []int, val int, begin, end int) int {
		if len(list) == 0 {
			return -1
		}

		delFirstPos := -1
		movePos := -1

		for i := begin; i < end; i++ {
			if list[i] == val && delFirstPos == -1 {
				delFirstPos = i
			} else {
				movePos = i
			}
		}

		if delFirstPos == -1 {
			return -1
		}

		if movePos == -1 {
			return delFirstPos
		}

		list[delFirstPos] = list[movePos]

		return delFirstPos
	}

	k := len(nums)
	delPos := 0
	for {
		delPos = dropOneElement(nums, val, delPos, k)
		if delPos == -1 {
			return k
		}
		k--
	}
}

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		act  []int
		k    int
	}{
		{nums: nil, val: 0, act: []int{}, k: 0},
		{nums: []int{1}, val: 1, act: []int{}, k: 0},
		{nums: []int{1, 1}, val: 1, act: []int{}, k: 0},
		{nums: []int{1, 2, 3, 4, 5}, val: -1, act: []int{1, 2, 3, 4, 5}, k: 5},
		{nums: []int{5, 4, 3, 2, 1}, val: -1, act: []int{1, 2, 3, 4, 5}, k: 5},
		{nums: []int{1, 2, 3, 4, 5}, val: 1, act: []int{2, 3, 4, 5}, k: 4},
		{nums: []int{1, 2, 3, 4, 5}, val: 4, act: []int{1, 2, 3, 5}, k: 4},
		{nums: []int{1, 2, 3, 4, 5}, val: 5, act: []int{1, 2, 3, 4}, k: 4},
		{nums: []int{1, 1, 2, 3, 4}, val: 1, act: []int{2, 3, 4}, k: 3},
		{nums: []int{1, 2, 3, 4, 2}, val: 2, act: []int{1, 3, 4}, k: 3},
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, act: []int{0, 0, 1, 3, 4}, k: 5},
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2, 2}, val: 2, act: []int{0, 0, 1, 3, 4}, k: 5},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.nums), func(t *testing.T) {
			nums := slices.Clone(test.nums)
			k := removeElement(nums, test.val)
			act := make([]int, 0, k)
			for i := 0; i < k; i++ {
				act = append(act, nums[i])
			}
			sort.Ints(act)

			require.Equal(t, test.k, k)
			require.Equal(t, test.act, act)
		})
	}
}
