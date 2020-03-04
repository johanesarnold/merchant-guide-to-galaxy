package main

import (
	"fmt"
	"prospace-test/helpers"
	"strings"
)

func main() {
	content := helpers.Read("test_input.txt")
	splittedInput := strings.Split(content, "\n")

	for _, input := range splittedInput {
		splittedRow := strings.Split(input, " is ")

		for _, row := range splittedRow {
			splittedChar := strings.Split(row, " ")

			fmt.Println("splitted: ", splittedChar[0])
		}
	}
	// numeral, err := helpers.NumeralConvert("23232")
	// helpers.FailOnError(err, "Failed to convert roman character")
	// fmt.Printf("%v", numeral)

}
