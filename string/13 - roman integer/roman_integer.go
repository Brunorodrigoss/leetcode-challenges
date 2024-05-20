package main

import (
	"fmt"
	"strings"
)

func romanToInt(input string) int {
	romanValueMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	inputSplited := strings.Split(input, "")
	result := 0
	previousValue, currentValue := 0, 0

	for index := 0; index < len(inputSplited); index++ {
		currentValue = romanValueMap[inputSplited[index]]

		if index == 0 {
			result += currentValue
		} else {
			if currentValue <= previousValue {
				result += currentValue
			} else {
				result += currentValue - (previousValue * 2)
			}
		}

		previousValue = currentValue
	}

	return result
}

func main() {
	fmt.Println(romanToInt("LXXIV"))
}
