package formatting_test

import (
	"testing"

	"github.com/joel-beck/sparplanrechner/pkg/formatting"
	"github.com/stretchr/testify/assert"
)

func TestFormatAmount_Integers(t *testing.T) {
	tests := []struct {
		name     string
		inputs   struct{ amount int }
		expected string
	}{
		{
			name:     "Zero integer",
			inputs:   struct{ amount int }{amount: int(0)},
			expected: "0,00",
		},
		{
			name:     "Above 1000",
			inputs:   struct{ amount int }{amount: 12340192},
			expected: "12.340.192,00",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := formatting.FormatAmount(test.inputs.amount)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestFormatAmount_Floats(t *testing.T) {
	tests := []struct {
		name     string
		inputs   struct{ amount float64 }
		expected string
	}{
		{
			name:     "Zero float",
			inputs:   struct{ amount float64 }{amount: float64(0.0)},
			expected: "0,00",
		},
		{
			name:     "Above 1000",
			inputs:   struct{ amount float64 }{amount: 1234.567},
			expected: "1.234,57",
		},
		{
			name:     "More than 2 decimal places",
			inputs:   struct{ amount float64 }{amount: 1234.56789},
			expected: "1.234,57",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := formatting.FormatAmount(test.inputs.amount)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestFormatAmounts_Integers(t *testing.T) {
	tests := []struct {
		name     string
		inputs   struct{ amounts []int }
		expected []string
	}{
		{
			name:     "Zero integer",
			inputs:   struct{ amounts []int }{amounts: []int{0}},
			expected: []string{"0,00"},
		},
		{
			name:     "Above 1000",
			inputs:   struct{ amounts []int }{amounts: []int{12340192}},
			expected: []string{"12.340.192,00"},
		},
		{
			name:     "Multiple integers",
			inputs:   struct{ amounts []int }{amounts: []int{12340192, 0, 1234}},
			expected: []string{"12.340.192,00", "0,00", "1.234,00"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := formatting.FormatAmounts(test.inputs.amounts)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestFormatAmounts_Floats(t *testing.T) {
	tests := []struct {
		name     string
		inputs   struct{ amounts []float64 }
		expected []string
	}{
		{
			name:     "Zero float",
			inputs:   struct{ amounts []float64 }{amounts: []float64{0.0}},
			expected: []string{"0,00"},
		},
		{
			name:     "Above 1000",
			inputs:   struct{ amounts []float64 }{amounts: []float64{1234.567}},
			expected: []string{"1.234,57"},
		},
		{
			name:     "More than 2 decimal places",
			inputs:   struct{ amounts []float64 }{amounts: []float64{1234.56789}},
			expected: []string{"1.234,57"},
		},
		{
			name:     "Multiple floats",
			inputs:   struct{ amounts []float64 }{amounts: []float64{1234.56789, 0.0, 1234.56789}},
			expected: []string{"1.234,57", "0,00", "1.234,57"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := formatting.FormatAmounts(test.inputs.amounts)
			assert.Equal(t, test.expected, actual)
		})
	}
}
