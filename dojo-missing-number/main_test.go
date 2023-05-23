package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	var want string = "Hello World!"
	output := Hello()
	if want != output {
		t.Fatalf(`Hello() = %s, expected %s`, output, want)
	}
}

func TestFoundMissingNumber(t *testing.T) {
	input := []int{0, 1}
	var want int = 2
	output := FoundMissingNumber(input)
	if want != output {
		t.Fatalf(`FoundMissingNumber(%v) = %d, expected %d`, input, output, want)
	}
}

func TestFoundMissingNumberWithTreePositionInputArray(t *testing.T) {
	input := []int{3, 0, 2}
	var want int = 1
	output := FoundMissingNumber(input)
	if want != output {
		t.Fatalf(`FoundMissingNumber(%v) = %d, expected %d`, input, output, want)
	}
}

func TestFoundMissingNumberWithNinePositionInputArray(t *testing.T) {
	input := []int{9, 6, 4, 2, 3, 8, 7, 0, 1}
	var want int = 5
	output := FoundMissingNumber(input)
	if want != output {
		t.Fatalf(`FoundMissingNumber(%v) = %d, expected %d`, input, output, want)
	}
}

func TestContaints(t *testing.T) {
	collection := []int{3, 0, 2}
	valueFromCollection := 2
	var want bool = true
	output := contains(collection, valueFromCollection)
	if want != output {
		t.Fatalf(`contains(%v, %d) = %t, expected %t`, collection, valueFromCollection, output, want)
	}
}

func TestContaintsShouldReturnFalse(t *testing.T) {
	collection := []int{3, 0, 2}
	valueFromCollection := 1
	var want bool = false
	output := contains(collection, valueFromCollection)
	if want != output {
		t.Fatalf(`contains(%v, %d) = %t, expected %t`, collection, valueFromCollection, output, want)
	}
}

func TestSecondCase(t *testing.T) {
	input := []int{2, 0}
	var want int = 1
	output := FoundMissingNumber2(input)
	if want != output {
		t.Fatalf(`FoundMissingNumber(%v) = %d, expected %d`, input, output, want)
	}
}

func TestThirdCase(t *testing.T) {
	input := []int{9, 6, 5, 2, 3, 8, 7, 0, 1}
	var want int = 4
	output := FoundMissingNumber2(input)
	if want != output {
		t.Fatalf(`FoundMissingNumber(%v) = %d, expected %d`, input, output, want)
	}
}
