package main

import (
	"errors"
	"os"
	"testing"
)

func TestArgsLength(t *testing.T) {
	expected := errors.New("Not enough arguments passed: Need at least 2. Test count and a test")
	os.Args = []string{"program", "200"}
	args := os.Args[1:]
	actual := validateArgs(args)

	if actual.Error() != expected.Error() {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestNumCasesLow(t *testing.T) {
	higherThanOne := errors.New("Number of test cases needs to be bigger than 1")
	os.Args = []string{"program", "0", "50"}
	args := os.Args[1:]
	actual := validateArgs(args)

	if actual == nil {
		t.Errorf("Test failed, expected: '%s', got: nil", higherThanOne)
	}

	if actual.Error() != higherThanOne.Error() {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", higherThanOne, actual)
	}
}

func TestNumCasesHigh(t *testing.T) {
	lowerThanHund := errors.New("Number of test cases must be lower than 100")
	os.Args = []string{"program", "200", "500"}
	args := os.Args[1:]
	actual := validateArgs(args)

	if actual.Error() != lowerThanHund.Error() {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", lowerThanHund, actual)
	}
}
