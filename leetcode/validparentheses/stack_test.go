package validparentheses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStack(t *testing.T) {
	s := stack[rune]{}
	require.Empty(t, s)

	s.push('a')
	require.EqualValues(t, []rune{'a'}, s)

	s.push('b')
	require.EqualValues(t, []rune{'a', 'b'}, s)

	require.Equal(t, 'b', s.pop())
	require.EqualValues(t, stack[rune]{'a'}, s)

	require.Equal(t, 'a', s.pop())
	require.EqualValues(t, stack[rune]{}, s)

	require.Panics(t, func() { s.pop() })
}
