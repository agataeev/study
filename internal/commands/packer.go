package commands

import (
	"fmt"
	"strconv"
	"unicode"
)

// Pack распаковывает строку
func Pack(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result string
	runes := []rune(input)
	count := 1

	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			count++
		} else {
			// Append the previous character and its count (if greater than 1)
			result += string(runes[i-1])
			if count > 1 {
				result += strconv.Itoa(count)
			}
			count = 1
		}
	}

	// Append the last character and its count
	result += string(runes[len(runes)-1])
	if count > 1 {
		result += strconv.Itoa(count)
	}

	return result
}

// Unpack распаковывает строку
func Unpack(input string) (string, error) {
	var result []rune
	runes := []rune(input)
	escapeMode := false

	for i := 0; i < len(runes); i++ {
		char := runes[i]

		if escapeMode {
			if char == '\\' || unicode.IsDigit(char) {
				result = append(result, char)
				escapeMode = false
			} else {
				return "", fmt.Errorf("invalid escape sequence at position %d", i)
			}
			continue
		}

		if char == '\\' {
			if i+1 < len(runes) && (runes[i+1] == '\\' || unicode.IsDigit(runes[i+1])) {
				escapeMode = true
				continue
			} else {
				return "", fmt.Errorf("invalid escape sequence at position %d", i)
			}
		}

		if unicode.IsLetter(char) {
			if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
				repeatCount, err := strconv.Atoi(string(runes[i+1]))
				if err != nil {
					return "", fmt.Errorf("invalid number at position %d", i+1)
				}

				for j := 0; j < repeatCount; j++ {
					result = append(result, char)
				}
				i++
			} else {
				result = append(result, char)
			}
		} else {
			result = append(result, char)
		}
	}

	return string(result), nil
}
