package helpers

import (
	"testing"
)

func TestNumeralConvert(t *testing.T) {
	romanArr := map[string]int{
		"I":         1,
		"II":        2,
		"III":       3,
		"IV":        4,
		"V":         5,
		"VI":        6,
		"VII":       7,
		"VIII":      8,
		"IX":        9,
		"X":         10,
		"XX":        20,
		"XXX":       30,
		"XL":        40,
		"L":         50,
		"LX":        60,
		"LXX":       70,
		"LXXX":      80,
		"XC":        90,
		"C":         100,
		"CC":        200,
		"CCC":       300,
		"CD":        400,
		"D":         500,
		"DC":        600,
		"DCC":       700,
		"DCCC":      800,
		"CM":        900,
		"M":         1000,
		"MM":        2000,
		"MMM":       3000,
		"MCMXC":     1990,
		"MMVIII":    2008,
		"MDCLXVI":   1666,
		"MMCDXXI":   2421,
		"CLX":       160,
		"CCVII":     207,
		"MIX":       1009,
		"MDCCLXXVI": 1776,
		"MCMLIV":    1954,
		"ZXCZXC":    -1,
		"1234":      -1,
		"0":         -1,
		"xl":        40,
		"m":         1000,
		"IIII":      -1,
		"XXXX":      -1,
		"CCCC":      -1,
		"MMMM":      -1,
		"DD":        -1,
		"LL":        -1,
		"VV":        -1,
		"IL":        -1,
		"IC":        -1,
		"ID":        -1,
		"IM":        -1,
		"XD":        -1,
		"XM":        -1,
		"VX":        -1,
		"VL":        -1,
		"VC":        -1,
		"VD":        -1,
		"VM":        -1,
		"LC":        -1,
		"LD":        -1,
		"LM":        -1,
		"DM":        -1,
		"A":         -1,
		"XXC":       -1,
	}

	for roman := range romanArr {
		got, err := NumeralConvert(roman)
		expected := romanArr[roman]

		if got != expected {
			t.Errorf("String: %v, Expected %d, Got %d (%s)", roman, expected, got, err)
		}
	}
}
