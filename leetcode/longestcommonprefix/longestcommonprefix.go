package longestcommonprefix

// LongestCommonPrefix longest common prefix in array words.
// Algo complexity O(n*m), where n length of array and m length of first word.
// Memory O(m).
func LongestCommonPrefix(words []string) string {
	if len(words) == 0 {
		return ""
	}

	firstWord := []rune(words[0])

	res := make([]rune, 0, len(firstWord))

	for i := 0; i < len(firstWord); i++ {
		curChar := firstWord[i]

		for _, w := range words[1:] {
			word := []rune(w)

			if len(word) == 0 {
				return ""
			}

			if i == len(word) {
				return string(res)
			}

			if word[i] != curChar {
				return string(res)
			}
		}

		res = append(res, curChar)
	}

	return string(res)
}
