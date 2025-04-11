package inttostr

var digits = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
}

func IntToSTr(n int) string {
	if n == 0 {
		return "0"
	}

	isNegative := n < 0

	var res string

	for n != 0 {
		digit := n % 10

		if digit < 0 {
			digit *= -1
		}

		res = digits[digit] + res
		n /= 10
	}

	if isNegative {
		res = "-" + res
	}

	return res
}
