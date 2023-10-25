package calculator

type InvestmentPlanRequest struct {
	StartCapital int     `json:"startCapital" form:"startCapital" query:"startCapital"`
	SavingsRate  int     `json:"savingsRate" form:"savingsRate" query:"savingsRate"`
	AnnualReturn float64 `json:"annualReturn" form:"annualReturn" query:"annualReturn"`
	Years        int     `json:"years" form:"years" query:"years"`
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

func (a *AnnualTotals) Sum() float64 {
	return Sum(*a)
}

// MonthlyPayments holds the individual saving rates for each month. They are not
// cumulative, nor do they include the starting capital.
type MonthlyPayments []int

func (e *MonthlyPayments) Sum(startCapital int) int {
	return Sum(*e) + startCapital
}

func (e *MonthlyPayments) AnnualCumulativeSum(startCapital int) []int {
	cumulativePayments := AnnualCumulativeSum(*e)
	for i := range cumulativePayments {
		cumulativePayments[i] += startCapital
	}
	return cumulativePayments
}

// MonthlyReturns holds the individual returns for each month. Just as with Payments, they are not cumulative.
type MonthlyReturns []float64

func (r *MonthlyReturns) Sum() float64 {
	return Sum(*r)
}

func (r *MonthlyReturns) AnnualCumulativeSum() []float64 {
	return AnnualCumulativeSum(*r)
}

type Amounts struct {
	CurrentTotal    float64
	AnnualTotals    AnnualTotals
	MonthlyPayments MonthlyPayments
	MonthlyReturns  MonthlyReturns
}

type Totals struct {
	totalAmount  float64
	totalPayment int
	totalReturns float64
}

func CalculateTotals(amounts Amounts, startCapital int) Totals {
	return Totals{
		totalAmount:  amounts.AnnualTotals.Sum(),
		totalPayment: amounts.MonthlyPayments.Sum(startCapital),
		totalReturns: amounts.MonthlyReturns.Sum(),
	}
}

type CumulativeAmounts struct {
	annualPayments []int
	annualReturns  []float64
}

func CalculateCumulativeAmounts(amounts Amounts, startCapital int) CumulativeAmounts {
	return CumulativeAmounts{
		annualPayments: amounts.MonthlyPayments.AnnualCumulativeSum(startCapital),
		annualReturns:  amounts.MonthlyReturns.AnnualCumulativeSum(),
	}
}
