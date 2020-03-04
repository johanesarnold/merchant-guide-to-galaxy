package helpers

import (
	"testing"
)

func TestNumeralConvert(t *testing.T) {
	romanArr := map[string]int{
		"MMCDXXI":   2421,
		"CLX":       160,
		"CCVII":     207,
		"MIX":       1009,
		"MDCCLXXVI": 1776,
		"MCMLIV":    1954,
		"ZXCZXC":    0,
		"1234":      0,
		"0":         0,
		"xl":        40,
		"m":         1000,
	}

	for roman := range romanArr {
		got, err := NumeralConvert(roman)
		expected := romanArr[roman]

		if err != nil {
			t.Errorf("String: %v, got err: %v", roman, err)
		} else if got != expected {
			t.Errorf("String: %v, Expected %d, Got %d", roman, expected, got)
		}
	}
}
