package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

const is = " is "
const space = " "
const wrong = "I have no idea what you are talking about"
const credits = "Credits"

func CalculateHowMuch(input string, romanMap map[string]string, valueMap map[string]float64) string {
	result := ""
	tempRoman := ""
	tempQuestion := ""
	if strings.Contains(input, " is ") {
		splitted := strings.Split(input, is)

		splittedValue := strings.Split(splitted[1], space)

		for index, value := range splittedValue {
			_, ok := romanMap[value]
			if ok {
				tempRoman += romanMap[value]
				tempQuestion += value
				if index == len(splittedValue)-2 {
					numeral, err := NumeralConvert(tempRoman)
					FailOnError(err, "Failed to convert roman character")
					tempQuestion += fmt.Sprintf(" is %d", numeral)
				} else {
					tempQuestion += " "
				}
			}

		}
		result = tempQuestion
	} else {
		result = wrong
	}
	return result
}

func CalculateHowMany(input string, romanMap map[string]string, valueMap map[string]float64) string {
	result := ""
	tempRoman := ""
	tempQuestion := ""

	if strings.Contains(input, " is ") {
		splitted := strings.Split(input, is)

		splittedValue := strings.Split(splitted[1], space)
		for index, value := range splittedValue {
			_, ok := romanMap[value]
			if ok {
				tempRoman += romanMap[value]
				tempQuestion += value
				tempQuestion += " "
			} else {
				_, ok := valueMap[value]
				if ok {
					if index == len(splittedValue)-2 {
						numeral, err := NumeralConvert(tempRoman)
						FailOnError(err, "Failed to convert roman character")
						finalNumeral := float64(numeral) * valueMap[value]

						tempQuestion += value
						tempQuestion += fmt.Sprintf(" is %d %s", int(finalNumeral), credits)
					} else {
						tempQuestion += " "
					}
				}
			}
		}
		result = tempQuestion
	} else {
		result = wrong
	}
	return result
}

func ReturnWrong() string {
	result := wrong
	return result
}

func StoreValueMap(input string, romanMap map[string]string) (string, float64) {
	splitted := strings.Split(input, is)
	splittedRoman := strings.Split(splitted[0], space)

	tempRoman := ""
	key := ""
	value := 0.00

	credits := strings.Split(splitted[1], space)
	creditValue, err := strconv.Atoi(credits[0])
	FailOnError(err, "Failed to convert string to int")

	for _, char := range splittedRoman {
		_, ok := romanMap[char]
		if ok {
			tempRoman += romanMap[char]
		} else {
			key = char
			numeral, err := NumeralConvert(tempRoman)
			FailOnError(err, "Failed to convert roman character")
			value = float64(creditValue) / float64(numeral)
		}

	}
	return key, value
}

func StoreRomanMap(input string) (string, string) {
	splitted := strings.Split(input, " is ")
	roman := strings.TrimSpace(splitted[1])

	key := splitted[0]
	value := roman

	return key, value
}
