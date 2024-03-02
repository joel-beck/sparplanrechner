package types

import "github.com/joel-beck/sparplanrechner/internal/formatting"

type FormattedIntermediateAmounts struct {
	AnnualTotals                    []string
	InflationDiscountedAnnualTotals []string
	AnnualPayments                  []string
	AnnualReturns                   []string
}

func FormatIntermediateAmounts(intermediateAmounts AnnualIntermediateAmounts) FormattedIntermediateAmounts {
	return FormattedIntermediateAmounts{
		AnnualTotals:                    formatting.FormatAmounts(intermediateAmounts.AnnualTotals),
		InflationDiscountedAnnualTotals: formatting.FormatAmounts(intermediateAmounts.InflationDiscountedAnnualTotals),
		AnnualPayments:                  formatting.FormatAmounts(intermediateAmounts.AnnualPayments),
		AnnualReturns:                   formatting.FormatAmounts(intermediateAmounts.AnnualReturns),
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

type ResponseData struct {
	Years               []int
	IntermediateAmounts FormattedIntermediateAmounts
	Totals              FormattedTotals
	Takeouts            FormattedTakeouts
}

func CollectResponseData(
	intermediateAmounts AnnualIntermediateAmounts,
	totals Totals,
	takeouts Takeouts,
	startCapital int) ResponseData {

	years := make([]int, len(intermediateAmounts.AnnualTotals))
	for i := range years {
		years[i] = i + 1
	}

	return ResponseData{
		Years:               years,
		IntermediateAmounts: FormatIntermediateAmounts(intermediateAmounts),
		Totals:              FormatTotals(totals),
		Takeouts:            FormatTakeouts(takeouts),
	}
}
