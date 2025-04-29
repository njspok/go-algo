package cuncurrentsorthead

import (
	"bufio"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

import (
	"io"
	"strings"
)

func ConcurrentSortHead(m int, files ...io.Reader) ([]string, error) {
	if m <= 0 {
		return nil, errors.New("count strings must positive integer")
	}
	if len(files) == 0 {
		return nil, errors.New("no files provided")
	}

	var scans []*bufio.Scanner
	for _, file := range files {
		scans = append(scans, bufio.NewScanner(file))
	}

	var res []string
	linesCounter := 0
	for {
		ready := 0
		for _, scan := range scans {
			if scan.Scan() {
				ready++
				linesCounter++
				res = append(res, scan.Text())
				if linesCounter == m {
					return res, nil
				}
			}
		}

		if ready == 0 {
			break
		}
	}

	if linesCounter < m {
		return res, errors.New("not enough lines")
	}

	return res, nil
}

func Test(t *testing.T) {
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
			count:  1,
			errMsg: "",
			files: []string{
				"aaa\nddd",
				"bbb\neee",
				"ccc\nfff",
			},
			result: []string{"aaa"},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var files []io.Reader
			for _, file := range tt.files {
				files = append(files, strings.NewReader(file))
			}

			act, err := ConcurrentSortHead(tt.count, files...)
			if tt.errMsg != "" {
				assert.Error(t, err, tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.result, act)
		})
	}
}
