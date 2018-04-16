package isin

import (
	"regexp"
)

var (
	index = map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
		'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14,
		'F': 15, 'G': 16, 'H': 17, 'I': 18, 'J': 19,
		'K': 20, 'L': 21, 'M': 22, 'N': 23, 'O': 24,
		'P': 25, 'Q': 26, 'R': 27, 'S': 28, 'T': 29,
		'U': 30, 'V': 31, 'W': 32, 'X': 33, 'Y': 34,
		'Z': 35,
		'a': 10, 'b': 11, 'c': 12, 'd': 13, 'e': 14,
		'f': 15, 'g': 16, 'h': 17, 'i': 18, 'j': 19,
		'k': 20, 'l': 21, 'm': 22, 'n': 23, 'o': 24,
		'p': 25, 'q': 26, 'r': 27, 's': 28, 't': 29,
		'u': 30, 'v': 31, 'w': 32, 'x': 33, 'y': 34,
		'z': 35,
	}

	isinFormat = regexp.MustCompile("^[A-Z]{2}[A-Z0-9]{9}[0-9]$")
)

// Validate function takes an isin and returns true if valid.
func Validate(isin string) bool {

	// Check if isin is formatted correctly
	if isinFormat.MatchString(isin) != true {
		return false
	}

	identifier := isin[:11]
	checkDigit := index[[]rune(isin[11:])[0]]

	var s []int
	for _, char := range identifier {
		switch {
		case index[char] > 10:
			s = append(s, index[char]/10)
			s = append(s, index[char]%10)
		default:
			s = append(s, index[char])

		}
	}

	sum := 0

	for i, digit := range s {
		switch i % 2 {
		case 0:
			digit := digit * 2
			sum += digit / 10
			sum += digit % 10

		default:
			sum += digit

		}
	}

	// Return false if check digit is incorrect.
	if ((10 - (sum % 10)) % 10) != int(checkDigit) {
		return false
	}

	return true
}
