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

func TestComputeAnnualTotals(t *testing.T) {
	tests := []struct {
		name     string
		inputs   types.UserInputs
		expected types.AnnualIntermediateTotals
	}{
		{
			name: "Zero Years",
			inputs: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           0,
				Years:                 0,
				AnnualReturnRate:      0,
				InflationRate:         0,
				InflationRateCheckbox: true,
				TakeoutRate:           0,
				TakeoutRateCheckbox:   true,
				Tax:                   0,
				TaxCheckbox:           true,
			},
			expected: types.AnnualIntermediateTotals{
				AnnualTotals:                    types.AnnualTotals{},
				InflationDiscountedAnnualTotals: types.AnnualTotals{},
				AnnualPayments:                  types.AnnualPayments{},
				AnnualReturns:                   types.AnnualReturns{},
			},
		},
		{
			name: "Only start capital",
			inputs: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           0,
				Years:                 1,
				AnnualReturnRate:      0,
				InflationRate:         0,
				InflationRateCheckbox: true,
				TakeoutRate:           0,
				TakeoutRateCheckbox:   true,
				Tax:                   0,
				TaxCheckbox:           true,
			},
			expected: types.AnnualIntermediateTotals{
				AnnualTotals:                    types.AnnualTotals{1000},
				InflationDiscountedAnnualTotals: types.AnnualTotals{1000},
				AnnualPayments:                  types.AnnualPayments{1000},
				AnnualReturns:                   types.AnnualReturns{0},
			},
		},
		{
			name: "Start capital and savings rate",
			inputs: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           100,
				Years:                 1,
				AnnualReturnRate:      0,
				InflationRate:         0,
				InflationRateCheckbox: true,
				TakeoutRate:           0,
				TakeoutRateCheckbox:   true,
				Tax:                   0,
				TaxCheckbox:           true,
			},
			expected: types.AnnualIntermediateTotals{
				AnnualTotals:                    types.AnnualTotals{2200},
				InflationDiscountedAnnualTotals: types.AnnualTotals{2200},
				AnnualPayments:                  types.AnnualPayments{2200},
				AnnualReturns:                   types.AnnualReturns{0},
			},
		},
		{
			name: "Start capital and inflation rate",
			inputs: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           0,
				Years:                 1,
				AnnualReturnRate:      0,
				InflationRate:         2,
				InflationRateCheckbox: true,
				TakeoutRate:           0,
				TakeoutRateCheckbox:   true,
				Tax:                   0,
				TaxCheckbox:           true,
			},
			expected: types.AnnualIntermediateTotals{
				AnnualTotals:                    types.AnnualTotals{1000},
				InflationDiscountedAnnualTotals: types.AnnualTotals{980},
				AnnualPayments:                  types.AnnualPayments{1000},
				AnnualReturns:                   types.AnnualReturns{0},
			},
		},
		{
			name: "Saving rate and inflation rate",
			inputs: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           100,
				Years:                 1,
				AnnualReturnRate:      0,
				InflationRate:         2,
				InflationRateCheckbox: true,
				TakeoutRate:           0,
				TakeoutRateCheckbox:   true,
				Tax:                   0,
				TaxCheckbox:           true,
			},
			expected: types.AnnualIntermediateTotals{
				AnnualTotals:                    types.AnnualTotals{2200},
				InflationDiscountedAnnualTotals: types.AnnualTotals{2156},
				AnnualPayments:                  types.AnnualPayments{2200},
				AnnualReturns:                   types.AnnualReturns{0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := calculator.ComputeAnnualTotals(tt.inputs)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestProcessCheckboxes(t *testing.T) {
	tests := []struct {
		name     string
		inputs   types.UserInputs
		expected types.UserInputs
	}{
		{
			name: "All checkboxes checked",
			inputs: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           100,
				Years:                 1,
				AnnualReturnRate:      5,
				InflationRate:         2,
				InflationRateCheckbox: true,
				TakeoutRate:           3,
				TakeoutRateCheckbox:   true,
				Tax:                   25,
				TaxCheckbox:           true,
			},
			expected: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           100,
				Years:                 1,
				AnnualReturnRate:      5,
				InflationRate:         2,
				InflationRateCheckbox: true,
				TakeoutRate:           3,
				TakeoutRateCheckbox:   true,
				Tax:                   25,
				TaxCheckbox:           true,
			},
		},
		{
			name: "Inflation rate checkbox disables inflation rate",
			inputs: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           100,
				Years:                 1,
				AnnualReturnRate:      5,
				InflationRate:         2,
				InflationRateCheckbox: false,
				TakeoutRate:           3,
				TakeoutRateCheckbox:   true,
				Tax:                   25,
				TaxCheckbox:           true,
			},
			expected: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           100,
				Years:                 1,
				AnnualReturnRate:      5,
				InflationRate:         0,
				InflationRateCheckbox: false,
				TakeoutRate:           3,
				TakeoutRateCheckbox:   true,
				Tax:                   25,
				TaxCheckbox:           true,
			},
		},
		{
			name: "Takeout rate checkbox disables takeout rate and tax",
			inputs: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           100,
				Years:                 1,
				AnnualReturnRate:      5,
				InflationRate:         2,
				InflationRateCheckbox: true,
				TakeoutRate:           3,
				TakeoutRateCheckbox:   false,
				Tax:                   25,
				TaxCheckbox:           true,
			},
			expected: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           100,
				Years:                 1,
				AnnualReturnRate:      5,
				InflationRate:         2,
				InflationRateCheckbox: true,
				TakeoutRate:           0,
				TakeoutRateCheckbox:   false,
				Tax:                   0,
				TaxCheckbox:           true,
			},
		},
		{
			name: "Tax checkbox disables tax",
			inputs: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           100,
				Years:                 1,
				AnnualReturnRate:      5,
				InflationRate:         2,
				InflationRateCheckbox: true,
				TakeoutRate:           3,
				TakeoutRateCheckbox:   true,
				Tax:                   25,
				TaxCheckbox:           false,
			},
			expected: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           100,
				Years:                 1,
				AnnualReturnRate:      5,
				InflationRate:         2,
				InflationRateCheckbox: true,
				TakeoutRate:           3,
				TakeoutRateCheckbox:   true,
				Tax:                   0,
				TaxCheckbox:           false,
			},
		},
		{
			name: "All checkboxes unchecked",
			inputs: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           100,
				Years:                 1,
				AnnualReturnRate:      5,
				InflationRate:         2,
				InflationRateCheckbox: false,
				TakeoutRate:           3,
				TakeoutRateCheckbox:   false,
				Tax:                   25,
				TaxCheckbox:           false,
			},
			expected: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           100,
				Years:                 1,
				AnnualReturnRate:      5,
				InflationRate:         0,
				InflationRateCheckbox: false,
				TakeoutRate:           0,
				TakeoutRateCheckbox:   false,
				Tax:                   0,
				TaxCheckbox:           false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := calculator.ProcessCheckboxes(tt.inputs)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestComputeAnnualTotals_ReturnRate(t *testing.T) {
	tests := []struct {
		name        string
		inputs      types.UserInputs
		lowerBounds struct {
			annualTotals                    float64
			inflationDiscountedAnnualTotals float64
			annualReturns                   float64
		}
	}{
		{
			name: "Annual interest compounding is lower bound to monthly interest compounding",
			inputs: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           0,
				Years:                 1,
				AnnualReturnRate:      5,
				InflationRate:         0,
				InflationRateCheckbox: true,
				TakeoutRate:           0,
				TakeoutRateCheckbox:   true,
				Tax:                   0,
				TaxCheckbox:           true,
			},
			lowerBounds: struct {
				annualTotals                    float64
				inflationDiscountedAnnualTotals float64
				annualReturns                   float64
			}{1050, 1050, 50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := calculator.ComputeAnnualTotals(tt.inputs)
			assert.Greater(t, actual.AnnualTotals[0], tt.lowerBounds.annualTotals)
			assert.Greater(t, actual.InflationDiscountedAnnualTotals[0], tt.lowerBounds.inflationDiscountedAnnualTotals)
			assert.Greater(t, actual.AnnualReturns[0], tt.lowerBounds.annualReturns)
		})
	}
}

func TestComputeAnnualTotals_MultipleYears(t *testing.T) {
	tests := []struct {
		name     string
		inputs   types.UserInputs
		expected types.AnnualIntermediateTotals
	}{
		{
			name: "Saving rate and inflation rate for two years",
			inputs: types.UserInputs{
				StartCapital:          1000,
				SavingsRate:           100,
				Years:                 2,
				AnnualReturnRate:      0,
				InflationRate:         2,
				InflationRateCheckbox: true,
				TakeoutRate:           0,
				TakeoutRateCheckbox:   true,
				Tax:                   0,
				TaxCheckbox:           true,
			},
			expected: types.AnnualIntermediateTotals{
				AnnualTotals:                    types.AnnualTotals{2200, 3400},
				InflationDiscountedAnnualTotals: types.AnnualTotals{2156, 3265.36},
				AnnualPayments:                  types.AnnualPayments{2200, 3400},
				AnnualReturns:                   types.AnnualReturns{0, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := calculator.ComputeAnnualTotals(tt.inputs)
			assert.Equal(t, tt.expected.AnnualTotals, actual.AnnualTotals)
			// account for rounding errors
			assert.InDeltaSlice(
				t,
				tt.expected.InflationDiscountedAnnualTotals,
				actual.InflationDiscountedAnnualTotals,
				0.01,
			)
			assert.Equal(t, tt.expected.AnnualPayments, actual.AnnualPayments)
			assert.Equal(t, tt.expected.AnnualReturns, actual.AnnualReturns)

		})
	}
}
