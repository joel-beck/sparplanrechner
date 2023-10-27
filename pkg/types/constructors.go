package types

import "github.com/joel-beck/sparplanrechner/pkg/calculator"

func MonthlyToAnnualTotals(
	annualTotals AnnualTotals,
	inflationDiscountedAnnualTotals AnnualTotals,
	monthlyIntermediates MonthlyIntermediateTotals,
	startCapital int,
) AnnualIntermediateTotals {
	return AnnualIntermediateTotals{
		AnnualTotals:                    annualTotals,
		InflationDiscountedAnnualTotals: inflationDiscountedAnnualTotals,
		AnnualPayments:                  monthlyIntermediates.MonthlyPayments.MonthlyToAnnual(startCapital),
		AnnualReturns:                   monthlyIntermediates.MonthlyReturns.MonthlyToAnnual(),
	}
}

func TotalsFromIntermediates(annualIntermediateTotals AnnualIntermediateTotals) Totals {
	return Totals{
		Total:                    annualIntermediateTotals.AnnualTotals.ComputeTotal(),
		InflationDiscountedTotal: annualIntermediateTotals.InflationDiscountedAnnualTotals.ComputeTotal(),
		Payments:                 annualIntermediateTotals.AnnualPayments.ComputeTotal(),
		Returns:                  annualIntermediateTotals.AnnualReturns.ComputeTotal(),
	}
}

func TakeoutsFromTotal(total float64, inputs UserInputs) Takeouts {
	t := Takeouts{}

	t.Annual.BeforeTax = total * inputs.TakeoutRate
	t.Annual.AfterTax = calculator.SubtractTax(t.Annual.BeforeTax, inputs.Tax)
	t.Annual.InflationDiscountedAfterTax = calculator.SubtractInflation(t.Annual.AfterTax, inputs.InflationRate, 1)

	t.Monthly.BeforeTax = t.Annual.BeforeTax / calculator.MonthsInYear
	t.Monthly.AfterTax = t.Annual.AfterTax / calculator.MonthsInYear
	t.Monthly.InflationDiscountedAfterTax = t.Annual.InflationDiscountedAfterTax / calculator.MonthsInYear

	return t
}
