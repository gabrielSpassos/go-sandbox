package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello())
}

func Hello() string {
	return "Hello World!"
}

func FoundMissingNumber2(input []int) int {
	size := len(input)
	// get max number by input lenght
	max_number := (size * (size + 1)) / 2

	// subtract input values from the max number
	for _, number := range input {
		max_number -= number
	}
	return max_number
}

func FoundMissingNumber(input []int) int {
	minValue := 0
	size := len(input)
	var expectedValues []int
	for i := minValue; i <= size; i++ {
		expectedValues = append(expectedValues, i)
	}

	for _, value := range expectedValues {
		valueIsOnExpectedValues := contains(input, value)
		if !valueIsOnExpectedValues {
			return value
		}
	}

	return -1
}

func contains(collection []int, value int) bool {
	for _, valueFromCollection := range collection {
		if valueFromCollection == value {
			return true
		}
	}
	return false
}
