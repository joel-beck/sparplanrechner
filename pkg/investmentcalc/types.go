package investmentcalc

type InvestmentPlanRequest struct {
	InitialCapital int     `json:"initialCapital" form:"initialCapital" query:"initialCapital"`
	SavingsRate    int     `json:"savingsRate" form:"savingsRate" query:"savingsRate"`
	AnnualReturn   float64 `json:"annualReturn" form:"annualReturn" query:"annualReturn"`
	Years          int     `json:"years" form:"years" query:"years"`
}
