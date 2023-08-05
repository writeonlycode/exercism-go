package isbn

import (
	"unicode/utf8"
)

var digits = map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
var check = map[rune]int{'X': 10}

func IsValidISBN(isbn string) bool {
	p := []int{}

	for _, v := range isbn {
		if d, ok := digits[v]; ok {
			p = append(p, d)
		} else if v != '-' && v != 'X' {
			return false
		}
	}

	lr, _ := utf8.DecodeLastRuneInString(isbn)

	if d, ok := check[lr]; ok {
		p = append(p, d)
	}

	if len(p) != 10 {
		return false
	}

	return (p[0]*10+p[1]*9+p[2]*8+p[3]*7+p[4]*6+p[5]*5+p[6]*4+p[7]*3+p[8]*2+p[9]*1)%11 == 0
}
