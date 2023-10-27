package types

type UserInputs struct {
	StartCapital     int     `json:"startCapital"     form:"startCapital"`
	SavingsRate      int     `json:"savingsRate"      form:"savingsRate"`
	Years            int     `json:"years"            form:"years"`
	AnnualReturnRate float64 `json:"annualReturnRate" form:"annualReturnRate"`
	InflationRate    float64 `json:"inflationRate"    form:"inflationRate"`
	TakeoutRate      float64 `json:"takeoutRate"      form:"takeoutRate"`
	Tax              float64 `json:"tax"              form:"tax"`
}
