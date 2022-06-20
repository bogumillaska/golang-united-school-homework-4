package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	trimed := strings.TrimSpace(input)
	if len(trimed) == 0 {
		return "", fmt.Errorf("wrong input: %w", errorEmptyInput)
	}

	var operands []string
	if strings.Contains(trimed, "+") {
		operands = strings.Split(trimed, "+")
	} else {
		lastMinus := strings.LastIndex(trimed, "-")
		if lastMinus == -1 {
			return "", fmt.Errorf("wrong input: %w", errorNotTwoOperands)
		} else {
			operands = []string{trimed[:lastMinus], trimed[lastMinus:]}
		}
	}

	if len(operands) != 2 {
		return "", fmt.Errorf("wrong input: %w", errorNotTwoOperands)
	}

	value1, err1 := strconv.Atoi(strings.TrimSpace(operands[0]))
	value2, err2 := strconv.Atoi(strings.TrimSpace(operands[1]))

	if err1 != nil {
		print(err1.Error())
		return "", fmt.Errorf(err1.Error())
	}

	if err2 != nil {
		print(err2.Error())
		return "", fmt.Errorf(err2.Error())
	}

	output = strconv.Itoa(value1 + value2)
	return output, nil
}
