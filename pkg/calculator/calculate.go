package calculator

import (
	"github.com/joel-beck/sparplanrechner/pkg/types"
)

func ZeroYearsOutput() types.AnnualIntermediateTotals {
	return types.AnnualIntermediateTotals{
		AnnualTotals:                    types.AnnualTotals{},
		InflationDiscountedAnnualTotals: types.AnnualTotals{},
		AnnualPayments:                  types.AnnualPayments{},
		AnnualReturns:                   types.AnnualReturns{},
	}
}

func initializeTotals() (types.MonthlyIntermediateTotals, types.AnnualIntermediateTotals) {
	return types.MonthlyIntermediateTotals{
			MonthlyPayments: types.MonthlyPayments{},
			MonthlyReturns:  types.MonthlyReturns{},
		},
		types.AnnualIntermediateTotals{
			AnnualTotals:                    types.AnnualTotals{},
			InflationDiscountedAnnualTotals: types.AnnualTotals{},
			AnnualPayments:                  types.AnnualPayments{},
			AnnualReturns:                   types.AnnualReturns{},
		}
}

func UpdateIntermediateTotals(
	inputs types.UserInputs,
	monthlyTotals *types.MonthlyIntermediateTotals,
	annualTotals *types.AnnualIntermediateTotals,
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
	monthlyIntermediates types.MonthlyIntermediateTotals,
	startCapital int,
) types.AnnualIntermediateTotals {
	return types.AnnualIntermediateTotals{
		AnnualTotals:                    annualTotals,
		InflationDiscountedAnnualTotals: inflationDiscountedAnnualTotals,
		AnnualPayments:                  monthlyIntermediates.MonthlyPayments.MonthlyToAnnual(startCapital),
		AnnualReturns:                   monthlyIntermediates.MonthlyReturns.MonthlyToAnnual(),
	}
}

func ComputeAnnualTotals(inputs types.UserInputs) types.AnnualIntermediateTotals {
	if inputs.Years == 0 {
		return ZeroYearsOutput()
	}

	totalMonths := inputs.Years * MonthsInYear
	monthlyReturn := CalculateMonthlyReturn(inputs.AnnualReturnRate)

	monthlyTotals, annualTotals := initializeTotals()
	currentTotal := float64(inputs.StartCapital)

	for month := 1; month <= totalMonths; month++ {
		UpdateIntermediateTotals(
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
