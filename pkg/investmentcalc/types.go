package investmentcalc

type InvestmentPlanRequest struct {
	InitialCapital float64 `json:"initialCapital" form:"initialCapital" query:"initialCapital"`
	AnnualReturn   float64 `json:"annualReturn" form:"annualReturn" query:"annualReturn"`
	Years          int     `json:"years" form:"years" query:"years"`
}
