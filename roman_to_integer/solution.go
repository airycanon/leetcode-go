package roman_to_integer

var single = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var double = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt(s string) int {
	runes := []rune(s)
	var (
		l     = len(runes)
		i     = 0
		sum   = 0
		value int
		exist bool
	)
	for i < l {

		if i+1 >= l {
			if value, exist = single[string(runes[i])]; exist {
				sum += value
			}
			return sum
		}

		doubleRoman := string(runes[i]) + string(runes[i+1])

		if value, exist = double[doubleRoman]; exist {
			sum += value
			i += 2
			continue
		}

		if value, exist = single[string(runes[i])]; exist {
			sum += value
			i += 1
			continue
		}
	}

	return sum
}
