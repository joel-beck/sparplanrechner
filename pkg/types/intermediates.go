package types

func LastElement[T any](a []T) T {
	return a[len(a)-1]
}

func Sum[T int | float64](a []T) T {
	var sum T
	for _, v := range a {
		sum += v
	}
	return sum
}

func AnnualCumulativeSum[T int | float64](a []T) []T {
	var sum T
	var annualCumulativeSums []T

	// only a single year of data
	if len(a) <= 12 {
		annualCumulativeSums = append(annualCumulativeSums, Sum(a))
		return annualCumulativeSums
	}

	for month, amount := range a {
		sum += amount

		if (month+1)%12 == 0 {
			annualCumulativeSums = append(annualCumulativeSums, sum)
		}
	}
	return annualCumulativeSums
}

// AnnualTotals holds the total amounts at the end of each year. Each value includes
// information about all previous years.
type AnnualTotals []float64

func (a AnnualTotals) ComputeTotal() float64 {
	return LastElement(a)
}

// MonthlyPayments holds the individual saving rates for each month. They are not
// cumulative, nor do they include the starting capital.
type MonthlyPayments []int

func (e *MonthlyPayments) MonthlyToAnnual(startCapital int) []int {
	cumulativePayments := AnnualCumulativeSum(*e)
	for i := range cumulativePayments {
		cumulativePayments[i] += startCapital
	}
	return cumulativePayments
}

type AnnualPayments []int

func (a AnnualPayments) ComputeTotal() int {
	return LastElement(a)
}

// MonthlyReturns holds the individual returns for each month. Just as with Payments,
// they are not cumulative.
type MonthlyReturns []float64

func (r *MonthlyReturns) MonthlyToAnnual() []float64 {
	return AnnualCumulativeSum(*r)
}

type AnnualReturns []float64

func (a AnnualReturns) ComputeTotal() float64 {
	return LastElement(a)
}

type MonthlyIntermediateTotals struct {
	MonthlyPayments MonthlyPayments
	MonthlyReturns  MonthlyReturns
}

type AnnualIntermediateTotals struct {
	AnnualTotals                    AnnualTotals
	InflationDiscountedAnnualTotals AnnualTotals
	AnnualPayments                  AnnualPayments
	AnnualReturns                   AnnualReturns
}
