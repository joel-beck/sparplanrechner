package calculator_test

import (
	"errors"
	"testing"

	"github.com/joel-beck/sparplanrechner/pkg/calculator"
	"github.com/stretchr/testify/assert"
)

func TestGrowthFactor(t *testing.T) {
	tests := []struct {
		name   string
		inputs struct {
			annualReturnRate float64
		}
		expected float64
	}{
		{
			name:     "Zero return rate",
			inputs:   struct{ annualReturnRate float64 }{annualReturnRate: 0},
			expected: 1,
		},
		{
			name:     "Positive return rate",
			inputs:   struct{ annualReturnRate float64 }{annualReturnRate: 5.0},
			expected: 1.05,
		},
		{
			name:     "Negative return rate",
			inputs:   struct{ annualReturnRate float64 }{annualReturnRate: -1.0},
			expected: 0.99,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := calculator.GrowthFactor(test.inputs.annualReturnRate)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestComputeTakeout_ValidInputs(t *testing.T) {
	tests := []struct {
		name   string
		inputs struct {
			total       float64
			takeoutRate float64
		}
		expected float64
	}{
		{
			name:     "Zero takeout rate",
			inputs:   struct{ total, takeoutRate float64 }{total: 100, takeoutRate: 0},
			expected: 0,
		},
		{
			name:     "Positive takeout rate",
			inputs:   struct{ total, takeoutRate float64 }{total: 100, takeoutRate: 1},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := calculator.ComputeTakeout(test.inputs.total, test.inputs.takeoutRate)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestComputeTakeout_InvalidInputs(t *testing.T) {
	tests := []struct {
		name   string
		inputs struct {
			total       float64
			takeoutRate float64
		}
		expected error
	}{
		{
			name:     "Negative takeout rate",
			inputs:   struct{ total, takeoutRate float64 }{total: 100, takeoutRate: -1},
			expected: errors.New("takeout rate must be non-negative"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := calculator.ComputeTakeout(test.inputs.total, test.inputs.takeoutRate)
			assert.Error(t, err)
			assert.Equal(t, test.expected, err)
		})
	}
}

func TestSubtractTax_ValidInputs(t *testing.T) {
	tests := []struct {
		name   string
		inputs struct {
			total   float64
			taxRate float64
		}
		expected float64
	}{
		{
			name:     "Zero tax rate",
			inputs:   struct{ total, taxRate float64 }{total: 100, taxRate: 0},
			expected: 100,
		},
		{
			name:     "Positive tax rate",
			inputs:   struct{ total, taxRate float64 }{total: 100, taxRate: 25},
			expected: 75,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := calculator.SubtractTax(test.inputs.total, test.inputs.taxRate)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSubtractTax_InvalidInputs(t *testing.T) {
	tests := []struct {
		name   string
		inputs struct {
			total   float64
			taxRate float64
		}
		expected error
	}{
		{
			name:     "Negative tax rate",
			inputs:   struct{ total, taxRate float64 }{total: 100, taxRate: -1},
			expected: errors.New("tax rate must be non-negative"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := calculator.SubtractTax(test.inputs.total, test.inputs.taxRate)
			assert.Error(t, err)
			assert.Equal(t, test.expected, err)
		})
	}
}

func TestInflationDiscountFactor(t *testing.T) {
	tests := []struct {
		name   string
		inputs struct {
			inflationRate float64
			years         float64
		}
		expected float64
	}{
		{
			name:     "Zero inflation rate",
			inputs:   struct{ inflationRate, years float64 }{inflationRate: 0, years: 1},
			expected: 1,
		},
		{
			name:     "Positive inflation rate",
			inputs:   struct{ inflationRate, years float64 }{inflationRate: 1, years: 1},
			expected: 0.99,
		},
		{
			name:     "Positive inflation rate two years",
			inputs:   struct{ inflationRate, years float64 }{inflationRate: 1, years: 2},
			expected: 0.9801,
		},
		{
			name:     "Negative inflation rate",
			inputs:   struct{ inflationRate, years float64 }{inflationRate: -1, years: 1},
			expected: 1.01,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := calculator.InflationDiscountFactor(test.inputs.inflationRate, test.inputs.years)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSubtractInflation(t *testing.T) {
	tests := []struct {
		name   string
		inputs struct {
			total         float64
			inflationRate float64
			years         float64
		}
		expected float64
	}{
		{
			name:     "Zero inflation rate",
			inputs:   struct{ total, inflationRate, years float64 }{total: 100, inflationRate: 0, years: 1},
			expected: 100,
		},
		{
			name:     "Positive inflation rate",
			inputs:   struct{ total, inflationRate, years float64 }{total: 100, inflationRate: 1, years: 1},
			expected: 99,
		},
		{
			name:     "Positive inflation rate two years",
			inputs:   struct{ total, inflationRate, years float64 }{total: 100, inflationRate: 1, years: 2},
			expected: 98.01,
		},
		{
			name:     "Negative inflation rate",
			inputs:   struct{ total, inflationRate, years float64 }{total: 100, inflationRate: -1, years: 1},
			expected: 101,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := calculator.SubtractInflation(test.inputs.total, test.inputs.inflationRate, test.inputs.years)
			assert.InDelta(t, test.expected, actual, 0.0001)
		})
	}
}

func TestIsFullYear(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "January",
			input:    1,
			expected: false,
		},
		{
			name:     "December",
			input:    12,
			expected: true,
		},
		{
			name:     "January next year",
			input:    13,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := calculator.IsFullYear(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestYearsFromMonth(t *testing.T) {
	tests := []struct {
		name     string
		input    struct{ month int }
		expected float64
	}{
		{
			name:     "zero month",
			input:    struct{ month int }{month: 0},
			expected: 0,
		},
		{
			name:     "multiple of 12",
			input:    struct{ month int }{month: 12},
			expected: 1,
		},
		{
			name:     "not multiple of 12",
			input:    struct{ month int }{month: 13},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := calculator.YearsFromMonth(test.input.month)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestComputeInflationDiscountedTotal(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			currentTotal, inflationRate float64
			month                       int
		}
		expected float64
	}{
		{
			name: "zero inflation rate",
			input: struct {
				currentTotal, inflationRate float64
				month                       int
			}{currentTotal: 100, inflationRate: 0, month: 10},
			expected: 100,
		},
		{
			name: "one year",
			input: struct {
				currentTotal, inflationRate float64
				month                       int
			}{currentTotal: 100, inflationRate: 1, month: 15},
			expected: 99,
		},
		{
			name: "two years",
			input: struct {
				currentTotal, inflationRate float64
				month                       int
			}{currentTotal: 100, inflationRate: 1, month: 27},
			expected: 98.01,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := calculator.ComputeInflationDiscountedTotal(
				test.input.currentTotal,
				test.input.inflationRate,
				test.input.month,
			)
			assert.InDelta(t, test.expected, actual, 0.0001)
		})
	}
}

func TestCalculateMonthlyReturn(t *testing.T) {
	tests := []struct {
		name   string
		inputs struct {
			annualReturnRate float64
		}
		expected float64
	}{
		{
			name:     "Zero return rate",
			inputs:   struct{ annualReturnRate float64 }{annualReturnRate: 0},
			expected: 0,
		},
		{
			name:     "Positive return rate",
			inputs:   struct{ annualReturnRate float64 }{annualReturnRate: 12.0},
			expected: 0.01,
		},
		{
			name:     "Negative return rate",
			inputs:   struct{ annualReturnRate float64 }{annualReturnRate: -12.0},
			expected: -0.01,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := calculator.CalculateMonthlyReturn(test.inputs.annualReturnRate)
			assert.Equal(t, test.expected, actual)
		})
	}
}
