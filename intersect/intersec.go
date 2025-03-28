package intersect

// Intersect two arrays
// Algo complexity O(n+m), O(n + min(n,m))
func Intersect(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return []int{}
	}

	cnt := make(map[int]int)

	for _, v := range a {
		cnt[v]++
	}

	res := make([]int, 0)

	for _, v := range b {
		if c, ok := cnt[v]; ok && c > 0 {
			cnt[v]--
			res = append(res, v)
		}
	}

	return res
}
