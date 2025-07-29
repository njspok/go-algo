package removeelements

import (
	"github.com/stretchr/testify/require"
	"slices"
	"sort"
	"strconv"
	"testing"
)

func removeElement(nums []int, val int) int {

	for i := range nums {
		if nums[i] != val {
			continue
		}

		nums = append(nums[:i], nums[i+1:]...)
		break
	}

	return len(nums)
}

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		act  []int
		k    int
	}{
		{nums: nil, val: 0, act: []int{}, k: 0},
		{nums: []int{1, 2, 3, 4, 5}, val: -1, act: []int{1, 2, 3, 4, 5}, k: 5},
		{nums: []int{5, 4, 3, 2, 1}, val: -1, act: []int{1, 2, 3, 4, 5}, k: 5},
		{nums: []int{1, 2, 3, 4, 5}, val: 1, act: []int{2, 3, 4, 5}, k: 4},
		{nums: []int{1, 2, 3, 4, 5}, val: 5, act: []int{1, 2, 3, 4}, k: 4},
	}
	for n, test := range tests {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
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
