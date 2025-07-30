package removeelements

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"slices"
	"sort"
	"testing"
)

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
		{nums: []int{1, 2}, val: 1, act: []int{2}, k: 1},
		{nums: []int{0, 1}, val: 1, act: []int{0}, k: 1},
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
		t.Run(fmt.Sprintf("%v %d", test.nums, test.val), func(t *testing.T) {
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
