package types

import (
	"github.com/joel-beck/sparplanrechner/pkg/calculator"
)

type Amounts struct {
	CurrentTotal                    float64
	AnnualTotals                    AnnualTotals
	InflationDiscountedAnnualTotals AnnualTotals
	MonthlyPayments                 MonthlyPayments
	MonthlyReturns                  MonthlyReturns
}

func UpdateAmounts(
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

	if calculator.IsFullYear(month) {
		amounts.AnnualTotals = append(amounts.AnnualTotals, amounts.CurrentTotal)

		inflationDiscountedTotal := calculator.ComputeInflationDiscountedTotal(
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
	totalMonths := inputs.Years * calculator.MonthsInYear
	monthlyReturn := calculator.CalculateMonthlyReturn(inputs.AnnualReturnRate)

	amounts := Amounts{}
	amounts.CurrentTotal = float64(inputs.StartCapital)

	for month := 1; month <= totalMonths; month++ {
		UpdateAmounts(&amounts, inputs, monthlyReturn, month)
	}

	return amounts
}
