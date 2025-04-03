package subsequence

// Subsequence return sub sequence based on seq.
// Algo complexity O(N^2) (more accuracy O(N^2/2)), memory O(N^2).
func Subsequence(seq []int) [][]int {
	if len(seq) == 0 {
		return [][]int{}
	}

	var res [][]int

	for card := 1; card <= len(seq); card++ {
		for i := 0; i <= len(seq)-card; i++ {
			el := seq[i : i+card]
			res = append(res, el)
		}
	}

	return res
}
