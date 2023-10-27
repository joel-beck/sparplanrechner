package calculator

import (
	"math"
)

const MonthsInYear = 12

// growthFactor calculates the growth factor for a given annual return rate. The annual return rate is provided in percent, not as a fraction.
func growthFactor(annualReturnRate float64) float64 {
	return 1 + (annualReturnRate / 100)
}

// inflationDiscountFactor calculates the discount factor for a given inflation rate and
// number of years. The inflation rate is provided in percent, not as a fraction.
// The underlying formula is (1 - inflationRate) ^ years where inflationRate is a
// fraction.
func inflationDiscountFactor(inflationRate float64, years float64) float64 {
	return math.Pow(1-(inflationRate/100), years)
}

func SubtractInflation(total float64, inflationRate float64, years float64) float64 {
	return total * inflationDiscountFactor(inflationRate, years)
}

func SubtractTax(total float64, taxRate float64) float64 {
	return total * (1 - taxRate)
}

func IsFullYear(month int) bool {
	return month%MonthsInYear == 0
}

func yearsFromMonth(month int) float64 {
	return float64(month) / float64(MonthsInYear)
}

func ComputeInflationDiscountedTotal(
	currentTotal float64,
	inflationRate float64,
	month int,
) float64 {
	years := yearsFromMonth(month)
	return SubtractInflation(currentTotal, inflationRate, years)
}

func CalculateMonthlyReturn(annualReturnRate float64) float64 {
	return (growthFactor(annualReturnRate) - 1) / MonthsInYear
}
