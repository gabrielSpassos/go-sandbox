package main

import (
	"testing"
)

func TestLookAndSay(t *testing.T) {
	var number int = 1
	var want int = 11
	output, err := LookAndSay(number)
	if want != output || err != nil {
		t.Fatalf(`LookAndSay("%d") = %d, expected %d`, number, output, want)
	}
}

func TestLookAndSayWithTwoDigitsNumber(t *testing.T) {
	var number int = 11
	var expected int = 21
	output, err := LookAndSay(number)
	if expected != output || err != nil {
		t.Fatalf(`LookAndSay("%d") = %d, expected %d`, number, output, expected)
	}
}

func TestLookAndSayWithMultipleDigitsNumber(t *testing.T) {
	var number int = 111221
	var expected int = 312211
	output, err := LookAndSay(number)
	if expected != output || err != nil {
		t.Fatalf(`LookAndSay("%d") = %d, expected %d`, number, output, expected)
	}
}

func TestLookAndSayMultipleTimes(t *testing.T) {
	result, err := LookAndSay(21)
	assert(t, result, err, 1211)

	result, err = LookAndSay(1211)
	assert(t, result, err, 111221)
}

func assert(t *testing.T, result int, err error, expect int) {
	if result != expect {
		t.Errorf("Error: %d != %d", expect, result)
	}
}
