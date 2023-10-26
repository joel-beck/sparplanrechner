package calculator

import (
	"fmt"
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
	fmt.Printf("Inflation rate: %f\n", inflationRate)
	return math.Pow(1-(inflationRate/100), years)
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
	// fmt.Printf("Inflation discounted total: %f\n", currentTotal*inflationDiscountFactor(inflationRate, years))
	return currentTotal * inflationDiscountFactor(inflationRate, years)
}

func calculateMonthlyReturn(annualReturnRate float64) float64 {
	return (growthFactor(annualReturnRate) - 1) / MonthsInYear
}

func updateAmounts(
	amounts *Amounts,
	savingsRate int,
	monthlyReturnRate float64,
	inflationRate float64,
	month int,
) {
	monthlyStartCapital := amounts.CurrentTotal + float64(savingsRate)
	monthlyReturns := monthlyStartCapital * monthlyReturnRate

	amounts.MonthlyPayments = append(amounts.MonthlyPayments, savingsRate)
	amounts.MonthlyReturns = append(amounts.MonthlyReturns, monthlyReturns)
	amounts.CurrentTotal = monthlyStartCapital + monthlyReturns

	if isFullYear(month) {
		amounts.AnnualTotals = append(amounts.AnnualTotals, amounts.CurrentTotal)

		fmt.Printf("Annual total after %d years: %s\n", month/MonthsInYear, formatAmount(amounts.CurrentTotal))

		inflationDiscountedTotal := computeInflationDiscountedTotal(
			amounts.CurrentTotal,
			inflationRate,
			month,
		)
		fmt.Printf(
			"Discounted annual total after %d years: %s\n",
			month/MonthsInYear,
			formatAmount(inflationDiscountedTotal),
		)

		amounts.InflationDiscountedAnnualTotals = append(
			amounts.InflationDiscountedAnnualTotals,
			inflationDiscountedTotal,
		)

	}

}

func CalculateAmounts(
	startCapital int,
	savingsRate int,
	annualReturnRate float64,
	years int,
	inflationRate float64,
) Amounts {
	totalMonths := years * MonthsInYear
	monthlyReturn := calculateMonthlyReturn(annualReturnRate)

	amounts := Amounts{}
	amounts.CurrentTotal = float64(startCapital)

	for month := 1; month <= totalMonths; month++ {
		updateAmounts(&amounts, savingsRate, monthlyReturn, inflationRate, month)
	}

	return amounts
}
