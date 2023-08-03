// Package bob defines a funtion to determine what Bob will reply to someone
// when they say something to him or ask him a question.
package bob

import (
	"strings"
	"unicode"
)

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func hasLetters(remark string) bool {
	for _, v := range remark {
		if unicode.IsLetter(v) {
			return true
		}
	}
	return false
}

func isYelling(remark string) bool {
	for _, v := range remark {
		if unicode.IsLetter(v) && !unicode.IsUpper(v) {
			return false
		}
	}

	return hasLetters(remark)
}

func isSilence(remark string) bool {
	for _, v := range remark {
		if !unicode.IsSpace(v) {
			return false
		}
	}

	return true
}

// Hey takes a remark and returns what Bob will reply.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	question := isQuestion(remark)
	yelling := isYelling(remark)

	switch {
	case isSilence(remark):
		return "Fine. Be that way!"
	case question && yelling:
		return "Calm down, I know what I'm doing!"
	case question:
		return "Sure."
	case yelling:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
