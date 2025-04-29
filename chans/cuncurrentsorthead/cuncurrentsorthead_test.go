package cuncurrentsorthead

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

import (
	"io"
	"strings"
)

func TestSortHead(t *testing.T) {
	tests := []struct {
		name   string
		count  int
		errMsg string
		files  []string
		result []string
	}{
		{
			name:   "no files",
			count:  4,
			files:  []string{},
			errMsg: "no files provided",
			result: nil,
		},
		{
			name:   "invalid count",
			count:  0,
			files:  []string{},
			errMsg: "count strings must positive integer",
			result: nil,
		},
		{
			name:  "empty files",
			count: 10,
			files: []string{
				"\n",
				"\n",
				"\n",
			},
			errMsg: "not enough lines",
			result: nil,
		},
		{
			name:   "success",
			count:  4,
			errMsg: "",
			files: []string{
				"aaa\nddd",
				"bbb\neee",
				"ccc\nfff",
			},
			result: []string{"aaa", "bbb", "ccc", "ddd"},
		},
		{
			name:   "success",
			count:  3,
			errMsg: "",
			files: []string{
				"aaa\nddd",
				"bbb\neee",
				"ccc\nfff",
			},
			result: []string{"aaa", "bbb", "ccc"},
		},
		{
			name:   "success",
			count:  6,
			errMsg: "",
			files: []string{
				"aaa\nddd",
				"bbb\neee",
				"ccc\nfff",
			},
			result: []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff"},
		},
		{
			name:   "success",
			count:  7,
			errMsg: "not enough lines",
			files: []string{
				"aaa\nddd",
				"bbb\neee",
				"ccc\nfff",
			},
			result: []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff"},
		},
		{
			name:   "success",
			count:  6,
			errMsg: "",
			files: []string{
				"ccc\nfff",
				"bbb\neee",
				"aaa\nddd",
			},
			result: []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff"},
		},
		{
			name:   "success",
			count:  6,
			errMsg: "not enough lines",
			files: []string{
				"ccc\n",
				"bbb\neee",
				"aaa\nddd",
			},
			result: []string{"aaa", "bbb", "ccc", "ddd", "eee"},
		},
		{
			name:   "success",
			count:  6,
			errMsg: "not enough lines",
			files: []string{
				"ccc\n",
				"\neee",
				"aaa\nddd",
			},
			result: []string{"aaa", "ccc", "ddd", "eee"},
		},
		{
			name:   "success",
			count:  100,
			errMsg: "not enough lines",
			files: []string{
				"\n\n\n\n",
				"\nbbb\n\n",
				"aaa\nccc",
			},
			result: []string{"aaa", "bbb", "ccc"},
		},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprintf("%d %s count %d", n, tt.name, tt.count), func(t *testing.T) {
			var files []io.Reader
			for _, file := range tt.files {
				files = append(files, strings.NewReader(file))
			}

			//act, err := SortHead(tt.count, files...)
			act, err := ConcurrencySortHead(tt.count, files...)

			if tt.errMsg != "" {
				assert.EqualError(t, err, tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.result, act)
		})
	}
}
