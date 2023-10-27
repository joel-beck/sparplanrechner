package types

import "github.com/joel-beck/sparplanrechner/pkg/formatting"

type FormattedIntermediateTotals struct {
	AnnualTotals                    []string
	InflationDiscountedAnnualTotals []string
	AnnualPayments                  []string
	AnnualReturns                   []string
}

func FormatIntermediateTotals(intermediateTotals AnnualIntermediateTotals) FormattedIntermediateTotals {
	return FormattedIntermediateTotals{
		AnnualTotals:                    formatting.FormatAmounts(intermediateTotals.AnnualTotals),
		InflationDiscountedAnnualTotals: formatting.FormatAmounts(intermediateTotals.InflationDiscountedAnnualTotals),
		AnnualPayments:                  formatting.FormatAmounts(intermediateTotals.AnnualPayments),
		AnnualReturns:                   formatting.FormatAmounts(intermediateTotals.AnnualReturns),
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
		Total:                    formatting.FormatAmount(totals.Total),
		InflationDiscountedTotal: formatting.FormatAmount(totals.InflationDiscountedTotal),
		Payments:                 formatting.FormatAmount(totals.Payments),
		Returns:                  formatting.FormatAmount(totals.Returns),
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
			BeforeTax:                   formatting.FormatAmount(takeouts.Monthly.BeforeTax),
			AfterTax:                    formatting.FormatAmount(takeouts.Monthly.AfterTax),
			InflationDiscountedAfterTax: formatting.FormatAmount(takeouts.Monthly.InflationDiscountedAfterTax),
		},
		Annual: FormattedAnnualTakeouts{
			BeforeTax:                   formatting.FormatAmount(takeouts.Annual.BeforeTax),
			AfterTax:                    formatting.FormatAmount(takeouts.Annual.AfterTax),
			InflationDiscountedAfterTax: formatting.FormatAmount(takeouts.Annual.InflationDiscountedAfterTax),
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
