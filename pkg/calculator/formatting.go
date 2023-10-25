package calculator

import (
	"fmt"
	"math"
	"strings"
)

func roundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}

func germanNumberFormat(f float64) string {
	str := fmt.Sprintf("%.2f", f)    // Convert float to string with 2 decimal places
	parts := strings.Split(str, ".") // Split integer and decimal parts
	integer := parts[0]
	decimal := parts[1]

	// Reverse the integer part string for easier processing
	reversed := ""
	for _, c := range integer {
		reversed = string(c) + reversed
	}

	// Insert periods every 3 digits
	withThousandSep := ""
	for i, c := range reversed {
		if i > 0 && i%3 == 0 {
			withThousandSep = "." + withThousandSep
		}
		withThousandSep = string(c) + withThousandSep
	}

	// Combine the integer part, comma, and the decimal part
	germanFormatted := withThousandSep + "," + decimal

	return germanFormatted
}

func formatAmount[T int | float64](amount T) string {
	return germanNumberFormat(roundToTwoDecimals(float64(amount)))
}

func formatAmounts[T int | float64](amounts []T) map[int]string {
	yearAmountMap := make(map[int]string, len(amounts))

	for i, amount := range amounts {
		// starts with Year 0
		yearAmountMap[i] = formatAmount(amount)
	}
	return yearAmountMap
}

func FormatTotalAmount(annualTotals AnnualTotals) string {
	return formatAmount(annualTotals.Sum())
}

func FormatAnnualTotals(annualTotals AnnualTotals) map[int]string {
	return formatAmounts(annualTotals)
}

func FormatTotalPayments(monthlyPayments MonthlyPayments, startCapital int) string {
	return formatAmount(float64(monthlyPayments.Sum(startCapital)))
}

func FormatAnnualPayments(monthlyPayments MonthlyPayments, startCapital int) map[int]string {
	return formatAmounts(monthlyPayments.AnnualCumulativeSum(startCapital))
}

func FormatTotalReturns(monthlyReturns MonthlyReturns) string {
	return formatAmount(monthlyReturns.Sum())
}

func FormatAnnualReturns(monthlyReturns MonthlyReturns) map[int]string {
	return formatAmounts(monthlyReturns.AnnualCumulativeSum())
}
