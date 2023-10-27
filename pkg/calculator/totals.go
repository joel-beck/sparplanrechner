package calculator

type Totals struct {
	Total                    float64
	InflationDiscountedTotal float64
	Payments                 int
	Returns                  float64
}

func TotalsFromIntermediates(annualIntermediateTotals AnnualIntermediateTotals) Totals {
	return Totals{
		Total:                    annualIntermediateTotals.AnnualTotals.ComputeTotal(),
		InflationDiscountedTotal: annualIntermediateTotals.InflationDiscountedAnnualTotals.ComputeTotal(),
		Payments:                 annualIntermediateTotals.AnnualPayments.ComputeTotal(),
		Returns:                  annualIntermediateTotals.AnnualReturns.ComputeTotal(),
	}
}
