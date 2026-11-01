package problem

import "unicode"

func isNumber(s string) bool {
	seenDigit := false
	seenExponent := false
	seenDot := false

	symbolMap := map[rune]bool{
		'+': true,
		'-': true,
	}
	exponentMap := map[rune]bool{
		'E': true,
		'e': true,
	}

	for i, r := range s {
		if unicode.IsDigit(r) {

			seenDigit = true

		} else if symbolMap[r] {

			if i != 0 {
				previousRune := rune(s[i-1])
				if !exponentMap[previousRune] {
					return false
				}
			}

		} else if exponentMap[r] {
			if seenExponent == true || seenDigit == false {
				return false
			}

			previousRune := rune(s[i-1])
			if symbolMap[previousRune] {
				return false
			}

			seenExponent = true
			seenDigit = false

		} else if r == '.' {

			if seenDot == true || seenExponent == true {
				return false
			}
			seenDot = true

		} else {
			return false
		}
	}

	return seenDigit
}

func isNumberLeetCodeVersion(s string) bool {
	seenDigit := false
	seenExponent := false
	seenDot := false
	for i := 0; i < len(s); i++ {
		curr := s[i]
		if '0' <= curr && curr <= '9' {
			seenDigit = true
		} else if curr == '+' || curr == '-' {
			if i > 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		} else if curr == 'e' || curr == 'E' {
			if seenExponent || !seenDigit {
				return false
			}
			seenExponent = true
			seenDigit = false
		} else if curr == '.' {
			if seenDot || seenExponent {
				return false
			}
			seenDot = true
		} else {
			return false
		}
	}
	return seenDigit
}
