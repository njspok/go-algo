package firstoccurrence

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 || len(needle) == 0 {
		return -1
	}

	for i := range haystack {
		if i+len(needle) > len(haystack) {
			return -1
		}
		if haystack[i:i+len(needle)] != needle {
			i += len(needle)
			continue
		}
		return i
	}

	return -1
}
