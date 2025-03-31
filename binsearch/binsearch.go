package binsearch

// BinSearch
// Algo complexity O(log N)
func BinSearch(sorted []int, search int) int {
	if len(sorted) == 0 {
		return -1
	}

	low := 0
	top := len(sorted) - 1

	for {
		if top <= low {
			if sorted[low] == search {
				return low
			}
			return -1
		}

		med := (low + top) / 2

		if search == sorted[med] {
			return med
		}

		if search < sorted[med] {
			top = med - 1
			continue
		}

		if search > sorted[med] {
			low = med + 1
			continue
		}

	}

}
