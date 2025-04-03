package subsequence

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSubsequence(t *testing.T) {
	tests := []struct {
		seq []int
		res [][]int
	}{
		{seq: nil, res: [][]int{}},
		{seq: []int{}, res: [][]int{}},
		{seq: []int{11}, res: [][]int{{11}}},
		{seq: []int{11, 22}, res: [][]int{{11}, {22}, {11, 22}}},
		{
			seq: []int{11, 22, 33},
			res: [][]int{
				{11},
				{22},
				{33},
				{11, 22},
				{22, 33},
				{11, 22, 33},
			},
		},
		{
			seq: []int{11, 22, 33, 44},
			res: [][]int{
				{11},
				{22},
				{33},
				{44},
				{11, 22},
				{22, 33},
				{33, 44},
				{11, 22, 33},
				{22, 33, 44},
				{11, 22, 33, 44},
			},
		},
		{
			seq: []int{11, 22, 33, 44, 55},
			res: [][]int{
				{11},
				{22},
				{33},
				{44},
				{55},
				{11, 22},
				{22, 33},
				{33, 44},
				{44, 55},
				{11, 22, 33},
				{22, 33, 44},
				{33, 44, 55},
				{11, 22, 33, 44},
				{22, 33, 44, 55},
				{11, 22, 33, 44, 55},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.seq), func(t *testing.T) {
			require.Equal(t, tt.res, Subsequence(tt.seq))
		})
	}
}
