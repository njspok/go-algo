package twosum

// TwoSum returns the first 2 indexes of the array whose elements are equal to target in total.
// Algo complexity O(N^2), memory O(1)
// Problems
// - sum big integers may overflow
func TwoSum(arr []int, target int) []int {
	if len(arr) < 2 {
		return []int{}
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
