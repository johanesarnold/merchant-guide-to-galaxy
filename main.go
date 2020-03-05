package main

import (
	"fmt"
	"prospace-test/helpers"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var tempValue int = 0
	var key string
	var tempRoman string
	var tempRomanValue int
	var result int = 0
	valueMap := make(map[string]int)
	romanMap := make(map[string]string)
	questionMap := make(map[string]int)

	content := helpers.Read("test_input.txt")
	splittedInput := strings.Split(content, "\n")

	for indexInput, input := range splittedInput {

		if strings.Contains(input, " is ") {
			splittedRow := strings.Split(input, " is ")
			reg, err := regexp.Compile("[^a-zA-Z0-9]+")
			helpers.FailOnError(err, "Failed to compile regex")
			if indexInput < 7 {
				if indexInput < 4 {
					roman := splittedRow[1]
					roman = reg.ReplaceAllString(roman, "")
					numeral, err := helpers.NumeralConvert(roman)
					helpers.FailOnError(err, "Failed to convert roman character")
					valueMap[splittedRow[0]] = numeral
					romanMap[splittedRow[0]] = roman
				} else {
					tempValue = 0
					tempRoman = ""
					credits := strings.Split(splittedRow[1], " ")
					creditValue, err := strconv.Atoi(credits[0])
					helpers.FailOnError(err, "Failed to convert string to int")

					splittedChar := strings.Split(splittedRow[0], " ")
					for _, char := range splittedChar {
						if valueMap[char] != 0 {
							tempRoman += romanMap[char]
						} else {
							key = char
							numeral, err := helpers.NumeralConvert(tempRoman)
							helpers.FailOnError(err, "Failed to convert roman character")
							tempValue = numeral
						}
					}
					valueMap[key] = creditValue / tempValue
				}
			} else {
				splittedValue := strings.Split(splittedRow[1], " ")
				result = 0
				tempRomanValue = 0
				tempRoman = ""
				for indexValue, value := range splittedValue {
					if strings.Contains(splittedRow[0], "much") {
						tempRoman += romanMap[value]
						if indexValue == len(splittedValue)-1 {
							numeral, err := helpers.NumeralConvert(tempRoman)
							helpers.FailOnError(err, "Failed to convert roman character")
							result = numeral
						}
					} else {
						_, ok := romanMap[value]
						if ok {
							tempRoman += romanMap[value]
						}

						if indexValue == len(splittedValue)-2 {
							numeral, err := helpers.NumeralConvert(tempRoman)
							helpers.FailOnError(err, "Failed to convert roman character")
							tempRomanValue = numeral
							result = valueMap[value] * tempRomanValue
						}
					}
				}

				question := strings.Replace(splittedRow[1], "?", "is ", -1)
				questionMap[question] = result
			}
		} else {
			questionMap["I have no idea what you are talking about"] = 0
		}
	}

	for key, question := range questionMap {
		key = strings.Replace(key, "\r", "", -1)
		if question != 0 {
			fmt.Printf("%v %v\n", key, question)
		} else {
			fmt.Printf("%v\n", key)
		}
	}
}
