package downloads

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestPropertyBased(t *testing.T) {
	for range 50 {
		// Arrange
		table := []struct {
			url   string
			msg   string
			err   error
			match string
		}{
			{
				url:   "https://example.com/a.xml",
				msg:   "downloaded data from https://example.com/a.xml",
				err:   errors.New("failed to download data from https://example.com/a.xml: timeout"),
				match: "",
			},
			{
				url:   "https://example.com/b.xml",
				msg:   "downloaded data from https://example.com/b.xml",
				err:   errors.New("failed to download data from https://example.com/b.xml: timeout"),
				match: "",
			},
			{
				url:   "https://example.com/c.xml",
				msg:   "downloaded data from https://example.com/c.xml",
				err:   errors.New("failed to download data from https://example.com/c.xml: timeout"),
				match: "",
			},
		}

		var urls []string
		for _, row := range table {
			urls = append(urls, row.url)
		}

		// Act
		msgs, err := download(urls)

		// Assert

		require.LessOrEqual(t, len(msgs), len(table))

		// mark all success results
		for _, msg := range msgs {
			for i := range table {
				if msg == table[i].msg {
					table[i].match = "success"
				}
			}
		}

		// mark all error results
		if err != nil {
			for i := range table {
				if table[i].match == "" {
					if strings.Contains(err.Error(), table[i].err.Error()) {
						table[i].match = "error"
					}
				}
			}
		}

		// all rows must me market
		for _, r := range table {
			require.NotEmpty(t, r.match, fmt.Sprintf("%v", r))
		}
	}
}
