package longestcommonprefix

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs   []string
		result string
	}{
		{[]string{}, ""},
		{[]string{"", "flow", "fire"}, ""},
		{[]string{"fork", "", "fire"}, ""},
		{[]string{"fork", "f", "fire"}, "f"},
		{[]string{"fork", "fox", "doctor"}, ""},
		{[]string{"fork", "fox", "fire"}, "f"},
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"flower", "flow", "flowen"}, "flow"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"вата", "ваня", "вантус"}, "ва"},
		{[]string{"вата", "вино", "вантус"}, "в"},
		{[]string{"вата", "", "вантус"}, ""},
		{[]string{"валера", "валентин", "валя"}, "вал"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.strs), func(t *testing.T) {
			require.Equal(t, tt.result, LongestCommonPrefix(tt.strs))
		})
	}
}
