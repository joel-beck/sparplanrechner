package calculator

type MonthlyTakeouts struct {
	BeforeTax                   float64
	AfterTax                    float64
	InflationDiscountedAfterTax float64
}

type AnnualTakeouts struct {
	BeforeTax                   float64
	AfterTax                    float64
	InflationDiscountedAfterTax float64
}

type Takeouts struct {
	Monthly MonthlyTakeouts
	Annual  AnnualTakeouts
}

func subtractTax(total float64, taxRate float64) float64 {
	return total * (1 - taxRate)
}

func TakeoutsFromTotal(total float64, inputs UserInputs) Takeouts {
	t := Takeouts{}

	t.Annual.BeforeTax = total * inputs.TakeoutRate
	t.Annual.AfterTax = subtractTax(t.Annual.BeforeTax, inputs.Tax)
	t.Annual.InflationDiscountedAfterTax = subtractInflation(t.Annual.AfterTax, inputs.InflationRate, 1)

	t.Monthly.BeforeTax = t.Annual.BeforeTax / MonthsInYear
	t.Monthly.AfterTax = t.Annual.AfterTax / MonthsInYear
	t.Monthly.InflationDiscountedAfterTax = t.Annual.InflationDiscountedAfterTax / MonthsInYear

	return t
}
