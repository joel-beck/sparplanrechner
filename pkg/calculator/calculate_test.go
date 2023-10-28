package calculator_test

import (
	"testing"

	"github.com/joel-beck/sparplanrechner/pkg/calculator"
	"github.com/joel-beck/sparplanrechner/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestMonthlyToAnnualTotals(t *testing.T) {
	tests := []struct {
		name   string
		inputs struct {
			annualTotals                    types.AnnualTotals
			inflationDiscountedAnnualTotals types.AnnualTotals
			monthlyIntermediates            types.MonthlyIntermediateTotals
			startCapital                    int
		}
		expected types.AnnualIntermediateTotals
	}{
		{
			name: "Empty inputs",
			inputs: struct {
				annualTotals                    types.AnnualTotals
				inflationDiscountedAnnualTotals types.AnnualTotals
				monthlyIntermediates            types.MonthlyIntermediateTotals
				startCapital                    int
			}{
				annualTotals:                    types.AnnualTotals{},
				inflationDiscountedAnnualTotals: types.AnnualTotals{},
				monthlyIntermediates: types.MonthlyIntermediateTotals{
					MonthlyPayments: types.MonthlyPayments{},
					MonthlyReturns:  types.MonthlyReturns{},
				},
				startCapital: 0,
			},
			expected: types.AnnualIntermediateTotals{
				AnnualTotals:                    types.AnnualTotals{},
				InflationDiscountedAnnualTotals: types.AnnualTotals{},
				AnnualPayments:                  types.AnnualPayments{0},
				AnnualReturns:                   types.AnnualReturns{0},
			},
		},
		{
			name: "Two years of inputs",
			inputs: struct {
				annualTotals                    types.AnnualTotals
				inflationDiscountedAnnualTotals types.AnnualTotals
				monthlyIntermediates            types.MonthlyIntermediateTotals
				startCapital                    int
			}{
				annualTotals:                    types.AnnualTotals{3000, 4000},
				inflationDiscountedAnnualTotals: types.AnnualTotals{2700, 3600},
				monthlyIntermediates: types.MonthlyIntermediateTotals{
					MonthlyPayments: types.MonthlyPayments{
						10,
						10,
						10,
						10,
						10,
						10,
						10,
						10,
						10,
						10,
						10,
						10,
						20,
						20,
						20,
						20,
						20,
						20,
						20,
						20,
						20,
						20,
						20,
						20,
					},
					MonthlyReturns: types.MonthlyReturns{
						5,
						5,
						5,
						5,
						5,
						5,
						5,
						5,
						5,
						5,
						5,
						5,
						10,
						10,
						10,
						10,
						10,
						10,
						10,
						10,
						10,
						10,
						10,
						10,
					}},
				startCapital: 200,
			},
			expected: types.AnnualIntermediateTotals{
				AnnualTotals:                    types.AnnualTotals{3000, 4000},
				InflationDiscountedAnnualTotals: types.AnnualTotals{2700, 3600},
				AnnualPayments:                  types.AnnualPayments{320, 560},
				AnnualReturns:                   types.AnnualReturns{60, 180},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := calculator.MonthlyToAnnualTotals(
				tt.inputs.annualTotals,
				tt.inputs.inflationDiscountedAnnualTotals,
				tt.inputs.monthlyIntermediates,
				tt.inputs.startCapital,
			)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
