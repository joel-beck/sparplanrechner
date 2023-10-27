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

func formatAmounts[T int | float64](amounts []T) []string {
	output := make([]string, len(amounts))

	for _, amount := range amounts {
		output = append(output, formatAmount(amount))
	}

	return output
}
