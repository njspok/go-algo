package factorial

import "math"

// RecursiveFactorial calc recursive factorial.
// Return positive number if ok and 0 if error.
func RecursiveFactorial(n uint) uint {
	if n == 0 {
		return 1
	}

	res := RecursiveFactorial(n - 1)

	if checkOverflowMult(n, res) {
		return 0
	}

	return n * res
}

// CycleFactorial calc cyclic factorial.
// Return positive number if ok and 0 if error.
func CycleFactorial(n uint) uint {
	if n == 0 {
		return 1
	}

	var res uint = 1

	for n > 1 {
		if checkOverflowMult(n, res) {
			return 0
		}

		res *= n
		n--
	}

	return res
}

func checkOverflowMult(a, b uint) bool {
	return math.MaxUint/a < b
}
