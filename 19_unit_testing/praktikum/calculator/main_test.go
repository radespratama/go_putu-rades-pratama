package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	testCase := Arguments{
		Command: "add",
		Domain:  1,
		Pair:    2,
	}
	assert.Equal(t, 3, testCase.Add())
}

func TestSubstraction(t *testing.T) {
	testCase := Arguments{
		Command: "substract",
		Domain:  1,
		Pair:    2,
	}
	assert.Equal(t, -1, testCase.Substract())
}

func TestMultiplication(t *testing.T) {
	testCase := Arguments{
		Command: "multiply",
		Domain:  1,
		Pair:    2,
	}
	assert.Equal(t, 2, testCase.Multiply())
}

func TestDivision(t *testing.T) {
	testCase := Arguments{
		Command: "divide",
		Domain:  1,
		Pair:    2,
	}
	assert.Equal(t, float32(0.5), testCase.Divide())
}

func TestCalculationWithAddition(t *testing.T) {
	testCase := Arguments{
		Command: "add",
		Domain:  1,
		Pair:    2,
	}
	assert.Equal(t, 3, testCase.Calcu())
}

func TestCalculationWithSubstraction(t *testing.T) {
	testCase := Arguments{
		Command: "substract",
		Domain:  1,
		Pair:    2,
	}
	assert.Equal(t, -1, testCase.Calcu())
}

func TestCalculationWithMultiplication(t *testing.T) {
	testCase := Arguments{
		Command: "multiply",
		Domain:  1,
		Pair:    2,
	}
	assert.Equal(t, 2, testCase.Calcu())
}

func TestCalculationWithDivision(t *testing.T) {
	testCase := Arguments{
		Command: "divide",
		Domain:  1,
		Pair:    2,
	}
	assert.Equal(t, float32(0.5), testCase.Calcu())
}

func TestCalculationUnimplemented(t *testing.T) {
	testCase := Arguments{
		Command: "power",
		Domain:  1,
		Pair:    2,
	}
	assert.Equal(t, "unimplemented", testCase.Calcu())
}

func TestMain(t *testing.T) {
	main()
}
