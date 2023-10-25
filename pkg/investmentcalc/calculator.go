package investmentcalc

import (
	"fmt"
	"math"
	"strings"
)

func CalculateAmounts(initialCapital float64, annualReturn float64, years int) []float64 {
	amounts := make([]float64, years)

	for i := 0; i < years; i++ {
		if i == 0 {
			amounts[i] = initialCapital * (1 + (annualReturn / 100))
			continue
		}
		amounts[i] = amounts[i-1] * (1 + (annualReturn / 100))
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
