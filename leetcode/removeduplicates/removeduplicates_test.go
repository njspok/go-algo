package removeduplicates

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func removeDuplicates(nums []int) int {

}

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
		k    int
	}{
		{nums: []int{}, want: []int{}, k: 0},
		{nums: []int{1, 1, 1, 2, 2, 2, 3}, want: []int{1, 2, 3}, k: 3},
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: []int{0, 1, 2, 3, 4}, k: 5},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.nums), func(t *testing.T) {
			k := removeDuplicates(test.nums)
			require.Equal(t, test.k, k)

			for i := 0; i < test.k; i++ {
				require.Equal(t, test.nums[i], test.want[i], test.nums)
			}
		})
	}
}
