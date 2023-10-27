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

func subtractInflation(total float64, inflationRate float64, years float64) float64 {
	return total * inflationDiscountFactor(inflationRate, years)
}

func isFullYear(month int) bool {
	return month%MonthsInYear == 0
}

func yearsFromMonth(month int) float64 {
	return float64(month) / float64(MonthsInYear)
}

func computeInflationDiscountedTotal(
	currentTotal float64,
	inflationRate float64,
	month int,
) float64 {
	years := yearsFromMonth(month)
	return subtractInflation(currentTotal, inflationRate, years)
}

func calculateMonthlyReturn(annualReturnRate float64) float64 {
	return (growthFactor(annualReturnRate) - 1) / MonthsInYear
}

func updateAmounts(
	amounts *Amounts,
	inputs UserInputs,
	monthlyReturnRate float64,
	month int,
) {
	monthlyStartCapital := amounts.CurrentTotal + float64(inputs.SavingsRate)
	monthlyReturns := monthlyStartCapital * monthlyReturnRate

	amounts.MonthlyPayments = append(amounts.MonthlyPayments, inputs.SavingsRate)
	amounts.MonthlyReturns = append(amounts.MonthlyReturns, monthlyReturns)
	amounts.CurrentTotal = monthlyStartCapital + monthlyReturns

	if isFullYear(month) {
		amounts.AnnualTotals = append(amounts.AnnualTotals, amounts.CurrentTotal)

		inflationDiscountedTotal := computeInflationDiscountedTotal(
			amounts.CurrentTotal,
			inputs.InflationRate,
			month,
		)
		amounts.InflationDiscountedAnnualTotals = append(
			amounts.InflationDiscountedAnnualTotals,
			inflationDiscountedTotal,
		)
	}
}

func CalculateAmounts(inputs UserInputs) Amounts {
	totalMonths := inputs.Years * MonthsInYear
	monthlyReturn := calculateMonthlyReturn(inputs.AnnualReturnRate)

	amounts := Amounts{}
	amounts.CurrentTotal = float64(inputs.StartCapital)

	for month := 1; month <= totalMonths; month++ {
		updateAmounts(&amounts, inputs, monthlyReturn, month)
	}

	return amounts
}
