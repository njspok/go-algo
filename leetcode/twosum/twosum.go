package twosum

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
