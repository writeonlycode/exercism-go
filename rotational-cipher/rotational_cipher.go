package rotationalcipher

import (
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	result := ""

	for _, v := range plain {
		switch {
		case !unicode.IsLetter(v):
			result += string(v)
		case unicode.IsUpper(v):
			char := rune(((int(v) - 65 + shiftKey) % 26) + 65)
			result += string(char)
		case unicode.IsLower(v):
			char := rune(((int(v) - 97 + shiftKey) % 26) + 97)
			result += string(char)

		}
	}

	return result
}
