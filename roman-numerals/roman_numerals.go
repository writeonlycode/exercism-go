package romannumerals

import "errors"

var ones = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
var tens = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var hundreds = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var thousands = []string{"", "M", "MM", "MMM"}

func ToRomanNumeral(input int) (string, error) {
	result := ""

	if input <= 0 || 3999 < input {
		return "", errors.New("Values greater than 3999 aren't supported!")
	}

	result += thousands[input/1000%10]
	result += hundreds[input/100%10]
	result += tens[input/10%10]
	result += ones[input%10]

	return result, nil
}
