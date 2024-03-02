package types

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
