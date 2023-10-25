package calculator

// growthFactor calculates the growth factor for a given annual return rate in percent.
func growthFactor(annualReturnRate float64) float64 {
	return 1 + (annualReturnRate / 100)
}

func isFullYear(month int) bool {
	return month%12 == 0
}

// calculateMonthlyReturn calculates the monthly return rate in percent for a given
// annual return rate in percent.
func calculateMonthlyReturn(annualReturnRate float64) float64 {
	return (growthFactor(annualReturnRate) - 1) / 12
}

func updateAmounts(amounts *Amounts, savingsRate int, monthlyReturnRate float64, month int) {
	// capital at beginning of the month
	monthlyStartCapital := amounts.CurrentTotal + float64(savingsRate)
	// monthly returns throughout this month
	monthlyReturns := monthlyStartCapital * monthlyReturnRate

	amounts.MonthlyPayments = append(amounts.MonthlyPayments, savingsRate)
	amounts.MonthlyReturns = append(amounts.MonthlyReturns, monthlyReturns)
	amounts.CurrentTotal = monthlyStartCapital + monthlyReturns

	if isFullYear(month) {
		amounts.AnnualTotals = append(amounts.AnnualTotals, amounts.CurrentTotal)
	}
}

func CalculateAmounts(startCapital int, savingsRate int, annualReturnRate float64, years int) Amounts {
	totalMonths := years * 12
	monthlyReturn := calculateMonthlyReturn(annualReturnRate)

	amounts := Amounts{}
	amounts.CurrentTotal = float64(startCapital)

	for month := 1; month <= totalMonths; month++ {
		updateAmounts(&amounts, savingsRate, monthlyReturn, month)
	}

	return amounts
}
