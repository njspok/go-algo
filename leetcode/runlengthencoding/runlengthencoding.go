package runlengthencoding

import (
	"strconv"
	"strings"
)

// algo complexity O(c^n)
func countAndSay(n int) string {
	if n <= 1 {
		return "1"
	}

	res := "1"
	for i := 2; i <= n; i++ {
		res = rle(res)
	}

	return res
}

// algo complexity O(len(str))
func rle(str string) string {
	if len(str) == 0 {
		return ""
	}

	res := strings.Builder{}

	char := str[0]
	counter := 0
	for i := 0; i < len(str); i++ {
		if char == str[i] {
			counter++
		} else {
			res.WriteString(strconv.Itoa(counter))
			res.WriteByte(char)

			counter = 1
			char = str[i]
		}
	}

	if counter > 0 {
		res.WriteString(strconv.Itoa(counter))
		res.WriteByte(char)
	}

	return res.String()
}
