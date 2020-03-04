package main

import (
	"fmt"
	"prospace-test/helpers"
	"strings"
)

func main() {
	content := helpers.Read("test_input.txt")
	splittedArr := strings.Split(content, "\n")

	for _, row := range splittedArr {
		fmt.Printf("\n%v", row)
	}
	numeral, err := helpers.NumeralConvert("23232")
	helpers.FailOnError(err, "Failed to convert roman character")
	fmt.Printf("%v", numeral)

}
