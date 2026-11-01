package problem

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlNum(s[i]) {
			i++
		}

		for i < j && !isAlNum(s[j]) {
			j--
		}

		if toLower(s[i]) != toLower(s[j]) {
			return false
		}

		i++
		j--
	}
	return true
}

func isAlNum(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func toLower(c byte) byte {
	return byte(unicode.ToLower(rune(c)))
}

func isPalindromeStackAttempt(s string) bool {
	stack := []byte{}

	for i := 0; i < len(s)/2; i++ {
		stack = append(stack, s[i])
	}

	odd := len(s)%2 != 0
	offset := len(s) / 2
	if odd {
		offset += 1
	}

	for i := offset; i < len(s); i++ {
		curr := s[i]
		pop := stack[len(stack)-1]

		if unicode.IsLetter(rune(curr)) {
			lowerCurr := byte(unicode.ToLower(rune(curr)))
			if lowerCurr != pop {
				return false
			}
		} else {
			if curr != pop {
				return false
			}
		}
		stack = stack[:len(stack)-1]
	}

	return true
}

func strip(s string) string {
	var result strings.Builder

	for i := 0; i < len(s); i++ {
		curr := s[i]
		if ('a' <= curr && curr <= 'z') || ('A' <= curr && curr <= 'Z') || ('0' <= curr && curr <= '9') {
			result.WriteByte(curr)
		}
	}

	return result.String()
}
