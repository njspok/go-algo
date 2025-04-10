package gcd

func RecursiveGCD(a, b uint) uint {
	if a == b {
		return a
	}

	if a < b {
		a, b = b, a
	}

	return RecursiveGCD(a-b, b)
}

func CycleGCD(a, b uint) uint {
	if a == b {
		return a
	}

	for {
		if a < b {
			a, b = b, a
		}

		diff := a - b
		if diff == 0 {
			return a
		}

		a = diff
	}
}
