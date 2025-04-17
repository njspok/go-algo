package sorting

// InsertionSort sort by insertions.
// Algo complexity O(N^2), O(1).
func InsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		x := a[i]
		j := i
		for j > 0 && x < a[j-1] {
			a[j] = a[j-1]
			j--
		}
		a[j] = x
	}
}
