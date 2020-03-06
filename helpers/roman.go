package helpers

import (
	"errors"
	"strings"
)

var symbolMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func NumeralConvert(roman string) (int, error) {

	var numeral int

	roman = strings.ToUpper(roman)

	for i := 0; i < len(roman); i++ {
		value := symbolMap[string(roman[i])]

		if value == 0 {
			return -1, errors.New("Invalid roman character")
		}

		next := i + 1
		valueNext := 0

		if next < len(roman) {
			valueNext = symbolMap[string(roman[next])]
			if value < valueNext {
				numeral += valueNext - value
				i++ //skip 1 iterate because get value from next
			} else {
				numeral += value
			}
		} else {
			numeral += value
		}
	}
	return numeral, nil
}

func GetSymbolMap() map[string]int {
	return symbolMap
}
