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

func TestIsHappyNumberShouldReturnTrue(t *testing.T) {
	var want bool = true
	var number int = 7
	isHappyNumber := IsHappyNumber(number)
	if want != isHappyNumber {
		t.Fatalf(`IsHappyNumber(%d) = %t, expected %t`, number, isHappyNumber, want)
	}
}

func TestIsHappyNumber(t *testing.T) {
	var want bool = true
	var number int = 19
	isHappyNumber := IsHappyNumber(number)
	if want != isHappyNumber {
		t.Fatalf(`IsHappyNumber(%d) = %t, expected %t`, number, isHappyNumber, want)
	}
}

func TestIsHappyNumberShouldReturnFalse(t *testing.T) {
	var want bool = false
	var number int = 4
	isHappyNumber := IsHappyNumber(number)
	if want != isHappyNumber {
		t.Fatalf(`IsHappyNumber(%d) = %t, expected %t`, number, isHappyNumber, want)
	}
}

func TestFloyd(t *testing.T) {
	input := []int{2, 4, 1, 5, 3, 6, 8, 7, 4}
	want := 4
	result := Floyd(input)
	if want != result {
		t.Fatalf(`Floyd(%v) = %d, expected %d`, input, result, want)
	}
}

func TestIsHappyNumber2ShouldReturnTrue(t *testing.T) {
	var want bool = true
	var number int = 7
	isHappyNumber := IsHappyNumber2(number)
	if want != isHappyNumber {
		t.Fatalf(`IsHappyNumber(%d) = %t, expected %t`, number, isHappyNumber, want)
	}
}

func TestIsHappyNumber2ShouldReturnFalse(t *testing.T) {
	var want bool = false
	var number int = 4
	isHappyNumber := IsHappyNumber2(number)
	if want != isHappyNumber {
		t.Fatalf(`IsHappyNumber(%d) = %t, expected %t`, number, isHappyNumber, want)
	}
}

func TestIsHappyNumber2(t *testing.T) {
	var want bool = true
	var number int = 19
	isHappyNumber := IsHappyNumber(number)
	if want != isHappyNumber {
		t.Fatalf(`IsHappyNumber(%d) = %t, expected %t`, number, isHappyNumber, want)
	}
}
