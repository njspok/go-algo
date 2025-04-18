package runlengthencoding

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		n   int
		res string
	}{
		{-1, "1"},
		{0, "1"},
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			require.Equal(t, tt.res, countAndSay(tt.n))
		})
	}
}

func TestRLE(t *testing.T) {
	tests := []struct {
		str    string
		result string
	}{
		{str: "", result: ""},
		{str: "1", result: "11"},
		{str: "2", result: "12"},
		{str: "3", result: "13"},
		{str: "11", result: "21"},
		{str: "22", result: "22"},
		{str: "111", result: "31"},
		{str: "1111", result: "41"},
		{str: "1234", result: "11121314"},
		{str: "1223334444", result: "11223344"},
		{str: "1122334411223344", result: "2122232421222324"},
		{str: "1111111111111", result: "131"},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			require.Equal(t, tt.result, rle(tt.str))
		})
	}
}
