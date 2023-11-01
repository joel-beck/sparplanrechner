package types

type UserInputs struct {
	StartCapital          int     `json:"startCapital"          form:"startCapital"`
	SavingsRate           int     `json:"savingsRate"           form:"savingsRate"`
	Years                 int     `json:"years"                 form:"years"`
	AnnualReturnRate      float64 `json:"annualReturnRate"      form:"annualReturnRate"`
	InflationRate         float64 `json:"inflationRate"         form:"inflationRate"`
	InflationRateCheckbox bool    `json:"inflationRateCheckbox" form:"inflationRateCheckbox"`
	TakeoutRate           float64 `json:"takeoutRate"           form:"takeoutRate"`
	TakeoutRateCheckbox   bool    `json:"takeoutRateCheckbox"   form:"takeoutRateCheckbox"`
	Tax                   float64 `json:"tax"                   form:"tax"`
	TaxCheckbox           bool    `json:"taxCheckbox"           form:"taxCheckbox"`
}
