package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Hello())
}

func Hello() string {
	return "Hello World!"
}

func IsHappyNumber(number int) bool {
	visitedNumbers := make(map[int]bool)
	result := 0

	for result == 0 {
		result = isHappyNumber(&number, visitedNumbers)
	}

	return result == 1
}

func isHappyNumber(number *int, visitedNumbers map[int]bool) int {
	fmt.Println("input number:", *number)
	if visitedNumbers[*number] {
		return -1
	}

	if *number == 1 {
		return 1
	}

	sumOfNumbersSquare := getSquareSum(*number)

	visitedNumbers[*number] = true
	fmt.Println(visitedNumbers)
	*number = sumOfNumbersSquare
	fmt.Println("*****************")

	return 0
}

func IsHappyNumber2(number int) bool {
	slow := number
	fast := number

	for ok := true; ok; ok = slow != fast {
		slow = getSquareSum(slow)

		fast = getSquareSum(getSquareSum(fast))
	}

	return slow == 1
}

func getSquareSum(number int) int {
	numberAsString := strconv.Itoa(number)
	splitedNumbers := strings.Split(numberAsString, "")
	sumOfNumbersSquare := 0

	for _, value := range splitedNumbers {
		intValue, _ := strconv.Atoi(value)
		squareValue := intValue * intValue
		sumOfNumbersSquare += squareValue
	}

	return sumOfNumbersSquare
}

func Floyd(numbers []int) int {
	n := len(numbers)
	if n <= 1 {
		return -1
	}

	slow := numbers[0]
	fast := numbers[numbers[0]]

	for fast != slow {
		slow = numbers[slow]
		fast = numbers[numbers[fast]]
	}

	fast = 0
	for slow != fast {
		slow = numbers[slow]
		fast = numbers[fast]
	}

	return slow
}
