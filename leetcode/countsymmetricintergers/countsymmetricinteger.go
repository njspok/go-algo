package countsymmetricintergers

import "strconv"

var digits = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func CountSymmetricIntegers(low int, high int) int {
	count := 0
	for n := low; n <= high; n++ {
		if isSymmetric(n) {
			count++
		}
	}
	return count
}

func isSymmetric(n int) bool {
	str := strconv.Itoa(n)
	if !isEven(len(str)) {
		return false
	}

	half := len(str) / 2

	left := str[:half]
	right := str[half:]

	return sumByDigit(left) == sumByDigit(right)
}

func isEven(n int) bool {
	return n%2 == 0
}

func sumByDigit(str string) int {
	sum := 0
	for _, sym := range str {
		d := digits[sym]
		sum += d
	}
	return sum
}
