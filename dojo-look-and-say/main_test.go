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
