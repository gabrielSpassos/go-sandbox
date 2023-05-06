package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	value, err := LookAndSay(226666655531111)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
}

func LookAndSay(number int) (int, error) {
	numberAsString := strconv.Itoa(number)

	slicedNumbers := splitStringIntoSlice(numberAsString)
	finalNumberAsString := formatRepeatedValues(slicedNumbers)

	finalNumber, err := strconv.Atoi(finalNumberAsString)
	if err != nil {
		return 0, errors.New("invalid number")
	}

	return finalNumber, nil
}

func splitStringIntoSlice(value string) []string {
	slice := strings.Split(value, "")
	return slice
}

func formatRepeatedValues(slicedNumbers []string) string {
	var finalValue = ""
	counter := 1

	for i, slicedNumber := range slicedNumbers {
		nextIndex := i + 1
		if nextIndex >= len(slicedNumbers) {
			finalValue = fmt.Sprint(finalValue, counter, slicedNumber)
			break
		}

		nextValue := slicedNumbers[nextIndex]
		if slicedNumber == nextValue {
			counter++
		} else {
			finalValue = fmt.Sprint(finalValue, counter, slicedNumber)
			counter = 1
		}
	}

	return finalValue
}
