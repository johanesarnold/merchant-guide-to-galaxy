package main

import (
	"fmt"
	"prospace-test/helpers"
	"strings"
)

const many = "how many"
const much = "how much"
const credits = "Credits"
const questionMark = "?"

func main() {
	valueMap := make(map[string]float64)
	romanMap := make(map[string]string)

	content := helpers.Read("test_input.txt")
	splittedInput := strings.Split(content, "\n")

	for _, input := range splittedInput {
		input = strings.TrimSpace(input)

		if strings.Contains(input, questionMark) {
			if strings.Contains(input, much) {
				result := helpers.CalculateHowMuch(input, romanMap, valueMap)
				fmt.Println(result)
			} else if strings.Contains(input, many) {
				result := helpers.CalculateHowMany(input, romanMap, valueMap)
				fmt.Println(result)
			} else {
				result := helpers.ReturnWrong()
				fmt.Println(result)
			}
		} else {
			if strings.Contains(input, credits) {
				key, value := helpers.StoreValueMap(input, romanMap)
				valueMap[key] = value
			} else {
				key, value := helpers.StoreRomanMap(input)
				if key == "" && value == "" {
					result := helpers.ReturnWrong()
					fmt.Println(result)
				} else {
					romanMap[key] = value
				}
			}
		}
	}
}
