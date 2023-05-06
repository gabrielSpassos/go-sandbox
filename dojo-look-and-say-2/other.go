package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := strings.Split("HELLO", "")
	fmt.Print(str[1])
}

func LookAndSay(number int) int {
	value := strconv.Itoa(number)

	counter := 1
	lastChar := string(value[0])

	var result strings.Builder

	for i, ch := range value {
		if i == 0 {
			continue
		}

		if lastChar == string(ch) {
			counter++
		} else {
			result.WriteString(strconv.Itoa(counter))
			result.WriteString(lastChar)
			counter = 1
		}
		lastChar = string(ch)
	}

	result.WriteString(strconv.Itoa(counter))
	result.WriteString(lastChar)

	final, _ := strconv.Atoi(result.String())
	return final
}

func LookAndSay2(number int) int {
	var valueString string = strconv.Itoa(number)
	value := strings.Split(valueString, "")
	var result string = ""

	i := 0
	for i < len(value) {
		count := 1
		for i+1 < len(value) && (value[i] == value[i+1]) {
			i += 1
			count += 1
		}
		result += fmt.Sprint(count, value[i])
		i += 1
	}

	final, _ := strconv.Atoi(result)
	return final
}
