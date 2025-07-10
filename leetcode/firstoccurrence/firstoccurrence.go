package firstoccurrence

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 {
		return -1
	}

	for i := range haystack {
		if i+len(needle) > len(haystack) {
			return -1
		}

		if d := diffPos(haystack[i:i+len(needle)], needle); d >= 0 {
			i += d + 1
			continue
		}
		return i
	}

	return -1
}

// >=0 position where str1 diff str2
// -1 if str1 == str2
// str1 must equal len str2
func diffPos(str1, str2 string) int {
	for i := range str1 {
		if str1[i] != str2[i] {
			return i
		}
	}
	return -1
}
