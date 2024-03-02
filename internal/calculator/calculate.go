package calculator

import (
	"github.com/joel-beck/sparplanrechner/internal/types"
)

func ZeroYearsOutput() types.AnnualIntermediateAmounts {
	return types.AnnualIntermediateAmounts{
		AnnualTotals:                    types.AnnualTotals{},
		InflationDiscountedAnnualTotals: types.AnnualTotals{},
		AnnualPayments:                  types.AnnualPayments{},
		AnnualReturns:                   types.AnnualReturns{},
	}
}

func initializeTotals() (types.MonthlyIntermediateAmounts, types.AnnualIntermediateAmounts) {
	return types.MonthlyIntermediateAmounts{
			MonthlyPayments: types.MonthlyPayments{},
			MonthlyReturns:  types.MonthlyReturns{},
		},
		types.AnnualIntermediateAmounts{
			AnnualTotals:                    types.AnnualTotals{},
			InflationDiscountedAnnualTotals: types.AnnualTotals{},
			AnnualPayments:                  types.AnnualPayments{},
			AnnualReturns:                   types.AnnualReturns{},
		}
}

func UpdateIntermediateAmounts(
	inputs types.UserInputs,
	monthlyTotals *types.MonthlyIntermediateAmounts,
	annualTotals *types.AnnualIntermediateAmounts,
	currentTotal *float64,
	monthlyReturnRate float64,
	month int,
) {
	monthlyStartCapital := *currentTotal + float64(inputs.SavingsRate)
	monthlyReturns := monthlyStartCapital * monthlyReturnRate

	monthlyTotals.MonthlyPayments = append(monthlyTotals.MonthlyPayments, inputs.SavingsRate)
	monthlyTotals.MonthlyReturns = append(monthlyTotals.MonthlyReturns, monthlyReturns)
	// assumes "monatliche Verzinsung", i.e. interest is compounded monthly
	*currentTotal = monthlyStartCapital + monthlyReturns

	if IsFullYear(month) {
		annualTotals.AnnualTotals = append(annualTotals.AnnualTotals, *currentTotal)

		inflationDiscountedTotal := ComputeInflationDiscountedTotal(
			*currentTotal,
			inputs.InflationRate,
			month,
		)
		annualTotals.InflationDiscountedAnnualTotals = append(
			annualTotals.InflationDiscountedAnnualTotals,
			inflationDiscountedTotal,
		)
	}
}

func MonthlyToAnnualTotals(
	annualTotals types.AnnualTotals,
	inflationDiscountedAnnualTotals types.AnnualTotals,
	monthlyIntermediates types.MonthlyIntermediateAmounts,
	startCapital int,
) types.AnnualIntermediateAmounts {
	return types.AnnualIntermediateAmounts{
		AnnualTotals:                    annualTotals,
		InflationDiscountedAnnualTotals: inflationDiscountedAnnualTotals,
		AnnualPayments:                  monthlyIntermediates.MonthlyPayments.MonthlyToAnnual(startCapital),
		AnnualReturns:                   monthlyIntermediates.MonthlyReturns.MonthlyToAnnual(),
	}
}

func ProcessCheckboxes(inputs types.UserInputs) types.UserInputs {
	if !inputs.InflationRateCheckbox {
		inputs.InflationRate = 0
	}

	if !inputs.TakeoutRateCheckbox {
		inputs.TakeoutRate = 0
		// tax is only relevant if takeout rate is non-zero
		inputs.Tax = 0
	}

	if !inputs.TaxCheckbox {
		inputs.Tax = 0
	}

	return inputs
}

func ComputeAnnualTotals(inputs types.UserInputs) types.AnnualIntermediateAmounts {
	if inputs.Years == 0 {
		return ZeroYearsOutput()
	}

	inputs = ProcessCheckboxes(inputs)

	totalMonths := inputs.Years * MonthsInYear
	monthlyReturn := CalculateMonthlyReturn(inputs.ReturnRate)

	monthlyTotals, annualTotals := initializeTotals()
	currentTotal := float64(inputs.StartCapital)

	for month := 1; month <= totalMonths; month++ {
		UpdateIntermediateAmounts(
			inputs,
			&monthlyTotals,
			&annualTotals,
			&currentTotal,
			monthlyReturn,
			month,
		)
	}

	return MonthlyToAnnualTotals(
		annualTotals.AnnualTotals,
		annualTotals.InflationDiscountedAnnualTotals,
		monthlyTotals,
		inputs.StartCapital,
	)
}
