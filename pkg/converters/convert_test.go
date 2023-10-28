package converters_test

import (
	"testing"

	"github.com/joel-beck/sparplanrechner/pkg/converters"
	"github.com/joel-beck/sparplanrechner/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestTotalsFromIntermediates(t *testing.T) {
	tests := []struct {
		name     string
		input    types.AnnualIntermediateTotals
		expected types.Totals
	}{
		{
			name: "Empty inputs",
			input: types.AnnualIntermediateTotals{
				AnnualTotals:                    types.AnnualTotals{},
				InflationDiscountedAnnualTotals: types.AnnualTotals{},
				AnnualPayments:                  types.AnnualPayments{},
				AnnualReturns:                   types.AnnualReturns{},
			},
			expected: types.Totals{
				Total:                    0,
				InflationDiscountedTotal: 0,
				Payments:                 0,
				Returns:                  0,
			},
		},
		{
			name: "Five years of inputs",
			input: types.AnnualIntermediateTotals{
				AnnualTotals:                    types.AnnualTotals{1000, 2000, 3000, 4000, 5000},
				InflationDiscountedAnnualTotals: types.AnnualTotals{900, 1800, 2700, 3600, 4500},
				AnnualPayments:                  types.AnnualPayments{100, 200, 300, 400, 500},
				AnnualReturns:                   types.AnnualReturns{10, 20, 30, 40, 50},
			},
			expected: types.Totals{
				Total:                    5000,
				InflationDiscountedTotal: 4500,
				Payments:                 500,
				Returns:                  50,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := converters.TotalsFromIntermediates(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestTakeoutsFromTotal(t *testing.T) {
	tests := []struct {
		name   string
		inputs struct {
			total float64
			user  types.UserInputs
		}
		expected types.Takeouts
	}{
		{
			name: "Empty inputs",
			inputs: struct {
				total float64
				user  types.UserInputs
			}{
				total: 0,
				user:  types.UserInputs{},
			},
			expected: types.Takeouts{
				Monthly: types.MonthlyTakeouts{
					BeforeTax:                   0,
					AfterTax:                    0,
					InflationDiscountedAfterTax: 0,
				},
				Annual: types.AnnualTakeouts{
					BeforeTax:                   0,
					AfterTax:                    0,
					InflationDiscountedAfterTax: 0,
				},
			},
		},
		{
			name: "Default inputs",
			inputs: struct {
				total float64
				user  types.UserInputs
			}{
				total: 10000,
				user: types.UserInputs{
					TakeoutRate:   3.0,
					Tax:           25.0,
					InflationRate: 2.0,
					Years:         1},
			},
			expected: types.Takeouts{
				Monthly: types.MonthlyTakeouts{
					BeforeTax:                   25,
					AfterTax:                    18.75,
					InflationDiscountedAfterTax: 18.375,
				},
				Annual: types.AnnualTakeouts{
					BeforeTax:                   300,
					AfterTax:                    225,
					InflationDiscountedAfterTax: 220.5,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := converters.TakeoutsFromTotal(tt.inputs.total, tt.inputs.user)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
