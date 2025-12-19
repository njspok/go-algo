package plusone

func plusOne(digits []int) []int {
	flag := false
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i] += 1
		}

		if flag {
			digits[i] += 1
			flag = false
		}

		if digits[i] == 10 {
			digits[i] = 0
			flag = true
		}
	}

	if flag {
		return append([]int{1}, digits...)
	}

	return digits
}
