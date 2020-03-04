package main

import (
	"fmt"
	"prospace-test/helpers"
	"strings"
)

func main() {
	content := helpers.ReadFile("test_input.txt")
	splittedArr := strings.Split(content, "\n")

	for _, row := range splittedArr {
		fmt.Printf("\n%v", row)
	}

}
