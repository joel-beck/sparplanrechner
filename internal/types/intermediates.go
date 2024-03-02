package types

import "errors"

func LastElement[T any](a []T) (T, error) {
	if len(a) == 0 {
		var zero T
		return zero, errors.New("cannot get last element of empty slice")
	}
	return a[len(a)-1], nil
}

func ComputeTotal[T int | float64](a []T) T {
	last, err := LastElement(a)
	// return total of 0, if slice is empty
	if err != nil {
		return 0
	}
	return last
}

func Sum[T int | float64](a []T) T {
	var sum T
	for _, v := range a {
		sum += v
	}
	return sum
}

func AnnualCumulativeSums[T int | float64](a []T) []T {
	var sum T
	var annualSums []T

	// only a single year of data
	if len(a) <= 12 {
		annualSums = append(annualSums, Sum(a))
		return annualSums
	}

	for month, amount := range a {
		sum += amount

		if (month+1)%12 == 0 {
			annualSums = append(annualSums, sum)
		}
	}
	return annualSums
}

// AnnualTotals holds the total amounts at the end of each year. Each value includes
// information about all previous years.
type AnnualTotals []float64

func (a AnnualTotals) ComputeTotal() float64 {
	return ComputeTotal(a)
}

// MonthlyPayments holds the individual saving rates for each month. They are not
// cumulative, nor do they include the starting capital.
type MonthlyPayments []int

func (e *MonthlyPayments) MonthlyToAnnual(startCapital int) []int {
	annualPayments := AnnualCumulativeSums(*e)
	for i := range annualPayments {
		annualPayments[i] += startCapital
	}
	return annualPayments
}

type AnnualPayments []int

func (a AnnualPayments) ComputeTotal() int {
	return ComputeTotal(a)
}

// MonthlyReturns holds the individual returns for each month. Just as with Payments,
// they are not cumulative.
type MonthlyReturns []float64

func (r *MonthlyReturns) MonthlyToAnnual() []float64 {
	return AnnualCumulativeSums(*r)
}

type AnnualReturns []float64

func (a AnnualReturns) ComputeTotal() float64 {
	return ComputeTotal(a)
}

type MonthlyIntermediateAmounts struct {
	MonthlyPayments MonthlyPayments
	MonthlyReturns  MonthlyReturns
}

type AnnualIntermediateAmounts struct {
	AnnualTotals                    AnnualTotals
	InflationDiscountedAnnualTotals AnnualTotals
	AnnualPayments                  AnnualPayments
	AnnualReturns                   AnnualReturns
}
