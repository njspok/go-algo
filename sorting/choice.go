package sorting

func ChoiceSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	for cur := 0; cur < len(arr); cur++ {
		sub := arr[cur:]

		pos, _ := min(sub)
		pos += cur

		if arr[pos] < arr[cur] {
			arr[cur], arr[pos] = arr[pos], arr[cur]
		}
	}
}

func min(arr []int) (pos int, min int) {
	if len(arr) == 0 {
		panic("empty array")
	}

	pos = 0
	min = arr[pos]
	for i, v := range arr {
		if v < min {
			pos = i
			min = v
		}
	}
	return
}
