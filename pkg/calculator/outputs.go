package calculator

type FormattedIntermediateTotals struct {
	AnnualTotals                    []string
	InflationDiscountedAnnualTotals []string
	AnnualPayments                  []string
	AnnualReturns                   []string
}

func FormatIntermediateTotals(intermediateTotals AnnualIntermediateTotals) FormattedIntermediateTotals {
	return FormattedIntermediateTotals{
		AnnualTotals:                    formatAmounts(intermediateTotals.AnnualTotals),
		InflationDiscountedAnnualTotals: formatAmounts(intermediateTotals.InflationDiscountedAnnualTotals),
		AnnualPayments:                  formatAmounts(intermediateTotals.AnnualPayments),
		AnnualReturns:                   formatAmounts(intermediateTotals.AnnualReturns),
	}
}

type FormattedTotals struct {
	Total                    string
	InflationDiscountedTotal string
	Payments                 string
	Returns                  string
}

func FormatTotals(totals Totals) FormattedTotals {
	return FormattedTotals{
		Total:                    formatAmount(totals.Total),
		InflationDiscountedTotal: formatAmount(totals.InflationDiscountedTotal),
		Payments:                 formatAmount(totals.Payments),
		Returns:                  formatAmount(totals.Returns),
	}
}

type FormattedMonthlyTakeouts struct {
	BeforeTax                   string
	AfterTax                    string
	InflationDiscountedAfterTax string
}

type FormattedAnnualTakeouts struct {
	BeforeTax                   string
	AfterTax                    string
	InflationDiscountedAfterTax string
}

type FormattedTakeouts struct {
	Monthly FormattedMonthlyTakeouts
	Annual  FormattedAnnualTakeouts
}

func FormatTakeouts(takeouts Takeouts) FormattedTakeouts {
	return FormattedTakeouts{
		Monthly: FormattedMonthlyTakeouts{
			BeforeTax:                   formatAmount(takeouts.Monthly.BeforeTax),
			AfterTax:                    formatAmount(takeouts.Monthly.AfterTax),
			InflationDiscountedAfterTax: formatAmount(takeouts.Monthly.InflationDiscountedAfterTax),
		},
		Annual: FormattedAnnualTakeouts{
			BeforeTax:                   formatAmount(takeouts.Annual.BeforeTax),
			AfterTax:                    formatAmount(takeouts.Annual.AfterTax),
			InflationDiscountedAfterTax: formatAmount(takeouts.Annual.InflationDiscountedAfterTax),
		},
	}
}

func CollectTemplateData(
	intermediateTotals AnnualIntermediateTotals,
	totals Totals,
	takeouts Takeouts,
	startCapital int) map[string]interface{} {
	years := make([]int, len(intermediateTotals.AnnualTotals))
	for i := range years {
		years[i] = i + 1
	}

	return map[string]interface{}{
		"Years":              years,
		"IntermediateTotals": FormatIntermediateTotals(intermediateTotals),
		"Totals":             FormatTotals(totals),
		"Takeouts":           FormatTakeouts(takeouts),
	}
}
