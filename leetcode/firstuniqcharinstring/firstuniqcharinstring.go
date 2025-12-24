package firstuniqcharinstring

// Complexity
// - time O(n)
// - memory O(n)
func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	type rec struct {
		pos   int
		count int
	}

	table := make(map[int32]rec)

	for pos, char := range s {
		if r, ok := table[char]; !ok {
			table[char] = rec{pos: pos, count: 1}
		} else {
			r.count++
			table[char] = r
		}
	}

	res := -1

	for _, r := range table {
		if r.count > 1 {
			continue
		}

		if res == -1 {
			res = r.pos
		} else {
			if res > r.pos {
				res = r.pos
			}
		}
	}

	return res
}

// Complexity
// - time O(n)
// - memory O(n)
func firstUniqChar3(s string) int {
	runes := []rune(s)
	table := make(map[rune]int)

	for _, r := range runes {
		table[r]++
	}

	for pos, r := range runes {
		if table[r] == 1 {
			return pos
		}
	}

	return -1
}

// better solution
// https://leetcode.com/problems/first-unique-character-in-a-string/solutions/7182220/go-simple-solution-100-by-ducnta-nq35/
//func firstUniqChar2(s string) int {
//	// Only 26 lowercase letters
//	var freq [26]int
//
//	// Count frequency of each character
//	for i := 0; i < len(s); i++ {
//		freq[s[i]-'a']++
//	}
//
//	// Find first character with count 1
//	for i := 0; i < len(s); i++ {
//		if freq[s[i]-'a'] == 1 {
//			return i
//		}
//	}
//
//	return -1
//}
