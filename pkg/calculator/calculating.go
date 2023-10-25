package calculator

const MonthsInYear = 12

func growthFactor(annualReturnRate float64) float64 {
	return 1 + (annualReturnRate / 100)
}

func isFullYear(month int) bool {
	return month%MonthsInYear == 0
}

func calculateMonthlyReturn(annualReturnRate float64) float64 {
	return (growthFactor(annualReturnRate) - 1) / MonthsInYear
}

func updateAmounts(amounts *Amounts, savingsRate int, monthlyReturnRate float64, month int) {
	monthlyStartCapital := amounts.CurrentTotal + float64(savingsRate)
	monthlyReturns := monthlyStartCapital * monthlyReturnRate

	amounts.MonthlyPayments = append(amounts.MonthlyPayments, savingsRate)
	amounts.MonthlyReturns = append(amounts.MonthlyReturns, monthlyReturns)
	amounts.CurrentTotal = monthlyStartCapital + monthlyReturns

	if isFullYear(month) {
		amounts.AnnualTotals = append(amounts.AnnualTotals, amounts.CurrentTotal)
	}
}

func CalculateAmounts(startCapital int, savingsRate int, annualReturnRate float64, years int) Amounts {
	totalMonths := years * MonthsInYear
	monthlyReturn := calculateMonthlyReturn(annualReturnRate)

	amounts := Amounts{}
	amounts.CurrentTotal = float64(startCapital)

	for month := 1; month <= totalMonths; month++ {
		updateAmounts(&amounts, savingsRate, monthlyReturn, month)
	}

	return amounts
}
