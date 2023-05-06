package main

import (
	"testing"
)

func TestLookAndSay(t *testing.T) {
	assert(t, 1, 11)
	assert(t, 11, 21)
	assert(t, 21, 1211)
	assert(t, 1211, 111221)
	assert(t, 111221, 312211)
}

func TestLookAndSay2(t *testing.T) {
	assert2(t, 1, 11)
	assert2(t, 11, 21)
	assert2(t, 21, 1211)
	assert2(t, 1211, 111221)
	assert2(t, 111221, 312211)
}

func assert(t *testing.T, value int, expect int) {
	result := LookAndSay(value)
	if result != expect {
		t.Errorf("Error: %d != %d", expect, result)
	}
}

func assert2(t *testing.T, value int, expect int) {
	result := LookAndSay2(value)
	if result != expect {
		t.Errorf("Error: %d != %d", expect, result)
	}
}
