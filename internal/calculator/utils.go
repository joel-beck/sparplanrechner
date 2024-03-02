package calculator

import (
	"errors"
	"math"
)

const MonthsInYear = 12

// GrowthFactor calculates the growth factor for a given annual return rate. The annual return rate is provided in percent, not as a fraction.
func GrowthFactor(returnRate float64) float64 {
	return 1 + (returnRate / 100)
}

func ComputeTakeout(total float64, takeoutRate float64) (float64, error) {
	if takeoutRate < 0 {
		return 0, errors.New("takeout rate must be non-negative")
	}

	return total * (takeoutRate / 100), nil
}

func SubtractTax(total float64, taxRate float64) (float64, error) {
	if taxRate < 0 {
		return 0, errors.New("tax rate must be non-negative")
	}

	return total * (1 - taxRate/100), nil
}

// InflationDiscountFactor calculates the discount factor for a given inflation rate and
// number of years. The inflation rate is provided in percent, not as a fraction.
// The underlying formula is (1 - inflationRate) ^ years where inflationRate is a
// fraction.
func InflationDiscountFactor(inflationRate float64, years float64) float64 {
	return math.Pow(1-inflationRate/100, years)
}

func SubtractInflation(total float64, inflationRate float64, years float64) float64 {
	return total * InflationDiscountFactor(inflationRate, years)
}

func IsFullYear(month int) bool {
	return month%MonthsInYear == 0
}

func YearsFromMonth(month int) float64 {
	return math.Floor(float64(month) / float64(MonthsInYear))
}

func ComputeInflationDiscountedTotal(
	currentTotal float64,
	inflationRate float64,
	month int,
) float64 {
	years := YearsFromMonth(month)
	return SubtractInflation(currentTotal, inflationRate, float64(years))
}

func CalculateMonthlyReturn(returnRate float64) float64 {
	return (returnRate / 100) / MonthsInYear
}
