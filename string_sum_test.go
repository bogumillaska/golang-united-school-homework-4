package string_sum

import (
	"errors"
	"strconv"
	"testing"
)

func TestStringSum_empty(t *testing.T) {
	got := "   "
	want := ""
	expectedError := errors.New("input is empty")

	result, err := StringSum(got)
	if errors.Is(err, expectedError) {
		t.Error("got %w, wanted %w\n", err.Error(), expectedError.Error())
	}

	if result != want {
		t.Error("Result should be empty")
	}

}
func TestStringSum(t *testing.T) {
	got := "3+5"
	want := "8"

	result, err := StringSum(got)
	if err != nil {
		t.Error("Should not return error")
	}

	if want != result {
		t.Errorf("get %v, wanted %v", got, want)
	}

}

func TestStringSum_withSpaces(t *testing.T) {
	got := " 3 + 5  "
	want := "8"

	result, err := StringSum(got)
	if err != nil {
		t.Error("Should not return error")
	}

	if want != result {
		t.Errorf("get %v, wanted %v", got, want)
	}

}

func TestStringSum_whiteSpace(t *testing.T) {
	got := "-3 -   5"
	want := "-8"

	result, err := StringSum(got)
	if err != nil {
		t.Error("Should not return error")
	}

	if want != result {
		t.Errorf("get %v, wanted %v", got, want)
	}

}

func TestStringSum_firstOperandNegative(t *testing.T) {
	got := " -3 + 5  "
	want := "2"

	result, err := StringSum(got)
	if err != nil {
		t.Error("Should not return error")
	}

	if want != result {
		t.Errorf("get %v, wanted %v", got, want)
	}

}

func TestStringSum_secondOperandNegative(t *testing.T) {
	got := "-3-5"
	want := "-8"

	result, err := StringSum(got)
	if err != nil {
		t.Error("Should not return error")
	}

	if want != result {
		t.Errorf("get %v, wanted %v", got, want)
	}

}

func TestStringSum_nonNumbersInFirstOperand(t *testing.T) {
	got := "-asdasd-5"
	want := ""

	result, err := StringSum(got)
	if errors.Is(err, &strconv.NumError{}) {
		t.Error("Should not return error")
	}

	if want != result {
		t.Errorf("get %v, wanted %v", got, want)
	}

}

func TestStringSum_nonNumbersInSeondOperand(t *testing.T) {
	got := "-3-asdasd"
	want := ""

	result, err := StringSum(got)
	if errors.Is(err, &strconv.NumError{}) {
		t.Error("Should not return error")
	}

	if want != result {
		t.Errorf("get %v, wanted %v", got, want)
	}

}

func TestStringSum_wrongNumerOfOperands(t *testing.T) {
	got := " 3 + 5 + 4"
	want := ""
	expectedError := errors.New("expecting two operands, but received more or less")

	result, err := StringSum(got)
	if errors.Is(err, expectedError) {
		t.Error("got %w, wanted %w\n", err.Error(), expectedError.Error())
	}

	if want != result {
		t.Errorf("get %v, wanted %v", got, want)
	}

}
