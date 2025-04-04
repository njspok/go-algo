package sorting

// BubbleSort bubble sort.
// Algo complexity O(N^2), memory O(1).
func BubbleSort(list []int) {
	if len(list) <= 1 {
		return
	}

	for last := len(list) - 1; last >= 1; last-- {
		isSorted := true
		for cur := 0; cur <= last-1; cur++ {
			next := cur + 1
			if list[cur] > list[next] {
				list[cur], list[next] = list[next], list[cur]
				isSorted = false
			}
		}
		if isSorted {
			return
		}
	}
}
