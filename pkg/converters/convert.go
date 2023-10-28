package converters

import (
	"github.com/joel-beck/sparplanrechner/pkg/calculator"
	"github.com/joel-beck/sparplanrechner/pkg/types"
)

func TotalsFromIntermediates(annualIntermediateTotals types.AnnualIntermediateTotals) types.Totals {
	return types.Totals{
		Total:                    annualIntermediateTotals.AnnualTotals.ComputeTotal(),
		InflationDiscountedTotal: annualIntermediateTotals.InflationDiscountedAnnualTotals.ComputeTotal(),
		Payments:                 annualIntermediateTotals.AnnualPayments.ComputeTotal(),
		Returns:                  annualIntermediateTotals.AnnualReturns.ComputeTotal(),
	}
}

func TakeoutsFromTotal(total float64, inputs types.UserInputs) types.Takeouts {
	t := types.Takeouts{}

	// Takeout rate from user input is in percent, not a fraction
	t.Annual.BeforeTax, _ = calculator.ComputeTakeout(total, inputs.TakeoutRate)

	t.Annual.AfterTax, _ = calculator.SubtractTax(t.Annual.BeforeTax, inputs.Tax)
	t.Annual.InflationDiscountedAfterTax = calculator.SubtractInflation(
		t.Annual.AfterTax,
		inputs.InflationRate,
		float64(inputs.Years),
	)

	t.Monthly.BeforeTax = t.Annual.BeforeTax / calculator.MonthsInYear
	t.Monthly.AfterTax = t.Annual.AfterTax / calculator.MonthsInYear
	t.Monthly.InflationDiscountedAfterTax = t.Annual.InflationDiscountedAfterTax / calculator.MonthsInYear

	return t
}
