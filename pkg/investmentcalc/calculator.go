package investmentcalc

import (
	"math"
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

func GetFinalAmount(amounts []float64) float64 {
	return roundToTwoDecimals(amounts[len(amounts)-1])
}

func GetIntermediateAmounts(amounts []float64) map[int]float64 {
	yearAmountMap := make(map[int]float64, len(amounts)-1)

	for i, amount := range amounts[:len(amounts)-1] {
		// starts with Year 1
		yearAmountMap[i+1] = roundToTwoDecimals(amount)
	}

	return yearAmountMap
}
