package twosum

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		res    []int
	}{
		{arr: []int{}, target: 0, res: []int{}},
		{arr: []int{11}, target: 0, res: []int{}},
		{arr: []int{3, 2}, target: 5, res: []int{0, 1}},
		{arr: []int{3, 2}, target: 1, res: []int{}},
		{arr: []int{3, 2, 4}, target: 6, res: []int{1, 2}},
		{arr: []int{2, 7, 11, 15}, target: 9, res: []int{0, 1}},
		{arr: []int{2, 7, 11, 15, 1}, target: 16, res: []int{3, 4}},
		{arr: []int{-2, 7, 11, +2, 1}, target: 0, res: []int{0, 3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("arr: %v target: %d", tt.arr, tt.target), func(t *testing.T) {
			require.Equal(t, tt.res, TwoSum(tt.arr, tt.target))
		})
	}
}
