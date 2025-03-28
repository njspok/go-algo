package intersec

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	require.Equal(t, []int{}, Intersec([]int{}, []int{}))
	require.Equal(t, []int{}, Intersec([]int{11}, []int{}))
	require.Equal(t, []int{}, Intersec([]int{}, []int{11}))
	require.Equal(t, []int{11}, Intersec([]int{11}, []int{11}))
}
