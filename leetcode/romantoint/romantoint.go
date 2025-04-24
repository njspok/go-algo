package romantoint

var roman = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var exceptions = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

// RomanToInt convert string roman notation number to integer.
// Algo complexity O(n), where n is length string, memory O(1)
func RomanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}

	var sum int

	for i := 0; i < len(s); i++ {

		if i+2 <= len(s) {
			if v, ok := exceptions[s[i:i+2]]; ok {
				sum += v
				i++
				continue
			}
		}

		if v, ok := roman[s[i]]; ok {
			sum += v
			continue
		}

		panic("invalid roman string")
	}

	return sum
}
