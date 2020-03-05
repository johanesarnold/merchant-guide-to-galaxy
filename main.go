package main

import (
	"fmt"
	"prospace-test/helpers"
	"strings"
)

const many = "how many"
const much = "how much"
const credits = "Credits"

func main() {
	valueMap := make(map[string]float64)
	romanMap := make(map[string]string)

	content := helpers.Read("test_input.txt")
	splittedInput := strings.Split(content, "\n")

	for _, input := range splittedInput {
		if strings.Contains(input, much) {
			result := helpers.CalculateHowMuch(input, romanMap, valueMap)
			fmt.Println(result)
		} else if strings.Contains(input, many) {
			result := helpers.CalculateHowMany(input, romanMap, valueMap)
			fmt.Println(result)
		} else {
			if !strings.Contains(input, credits) {
				key, value := helpers.StoreRomanMap(input)
				romanMap[key] = value
			} else {
				key, value := helpers.StoreValueMap(input, romanMap)
				valueMap[key] = value
			}
		}
	}
}
