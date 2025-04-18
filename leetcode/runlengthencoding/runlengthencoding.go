package runlengthencoding

import "strconv"

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

func rle(str string) string {
	if len(str) == 0 {
		return ""
	}

	var res string
	char := str[0]
	counter := 0
	for i := 0; i < len(str); i++ {
		if char == str[i] {
			counter++
		} else {
			res += strconv.Itoa(counter) + string(char)
			counter = 1
			char = str[i]
		}
	}

	if counter > 0 {
		res += strconv.Itoa(counter) + string(char)
	}

	return res
}
