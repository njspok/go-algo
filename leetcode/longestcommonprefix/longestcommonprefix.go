package longestcommonprefix

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var res string

	firstWord := strs[0]

	for i := 0; i < len(firstWord); i++ {
		curChar := firstWord[i]

		for _, str := range strs[1:] {
			if len(str) == 0 {
				return ""
			}

			if i == len(str) {
				return res
			}

			if str[i] != curChar {
				return res
			}
		}

		res += string(curChar)
	}

	return res
}
