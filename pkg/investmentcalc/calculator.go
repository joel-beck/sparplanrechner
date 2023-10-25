package investmentcalc

import (
	"fmt"
	"math"
	"strings"
)

func growthFactor(annualReturn float64) float64 {
	return 1 + (annualReturn / 100)
}

func isFullYear(month int) bool {
	return month%12 == 0
}

func calculateMonthlyReturn(annualReturn float64) float64 {
	return (growthFactor(annualReturn) - 1) / 12
}

func updateCurrentTotal(currentTotal float64, savingsRate int, monthlyReturn float64) float64 {
	// capital at beginning of the month
	startCapital := currentTotal + float64(savingsRate)
	// returns throughout this month
	returns := startCapital * monthlyReturn

	return startCapital + returns
}

func appendYearlyAmount(month int, currentTotal float64, amounts *[]float64) {
	if isFullYear(month) {
		*amounts = append(*amounts, currentTotal)
	}
}

func CalculateAmounts(initialCapital int, savingsRate int, annualReturn float64, years int) []float64 {
	totalMonths := years * 12
	monthlyReturn := calculateMonthlyReturn(annualReturn)

	amounts := []float64{}
	currentTotal := float64(initialCapital)

	for month := 1; month <= totalMonths; month++ {
		currentTotal = updateCurrentTotal(currentTotal, savingsRate, monthlyReturn)
		appendYearlyAmount(month, currentTotal, &amounts)
	}

	return amounts
}

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

func GetFinalAmount(amounts []float64) string {
	return germanNumberFormat(roundToTwoDecimals(amounts[len(amounts)-1]))
}

func GetIntermediateAmounts(amounts []float64) map[int]string {
	yearAmountMap := make(map[int]string, len(amounts)-1)

	for i, amount := range amounts[:len(amounts)-1] {
		// starts with Year 1
		yearAmountMap[i+1] = germanNumberFormat(roundToTwoDecimals(amount))
	}

	return yearAmountMap
}
