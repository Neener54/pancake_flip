package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func validateArgs(args []string) error {
	if len(args) <= 1 {
		return errors.New("Not enough arguments passed: Need at least 2. Test count and a test")
	}

	length, err := strconv.Atoi(args[0])

	if err != nil {
		return err
	}

	if length <= 1 {
		return errors.New("Number of test cases needs to be bigger than 1")
	}

	if length >= 100 {
		return errors.New("Number of test cases must be lower than 100")
	}
	return nil
}
func flip(numPancakes int, pancakes *string) {
	arr := []byte(*pancakes)
	for i := 0; i < numPancakes; i++ {
		if arr[i] == '+' {
			arr[i] = '-'
		} else {
			arr[i] = '+'
		}
	}
	*pancakes = string(arr)
}
func handleCase(c string) int {
	p := 0
	numFlips := 0
	for loop := 0; loop < 100; loop++ {
		loop++

		for i := p; i < len(c); i++ {
			if c[i] != c[0] {
				break
			}
			p++
		}

		if p == len(c) {
			break
		}

		numFlips++
		flip(p, &c)
	}
	return numFlips
}

func flipPancakes(args []string) int {
	testCases := args[1:]

	err := validateArgs(args)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return 1
	}

	for caseIndex, testCase := range testCases {
		fmt.Printf("Case #%d: %d\n", caseIndex+1, handleCase(testCase))
	}
	return 0
}

func main() {
	args := os.Args[1:]
	flipPancakes(args)
}
