package types

import (
	"strconv"
	"strings"
)

func NumberInputToInt(s string) int {
	i, err := strconv.Atoi(strings.Replace(s, ".", "", -1))

	if err != nil {
		return 0
	}

	return i
}

func CheckboxInputToBool(s string) bool {
	return s == "on"
}

type RawUserInputs struct {
	StartCapital          string  `json:"startCapital"          form:"startCapital"`
	SavingsRate           string  `json:"savingsRate"           form:"savingsRate"`
	Years                 int     `json:"years"                 form:"years"`
	AnnualReturnRate      float64 `json:"annualReturnRate"      form:"annualReturnRate"`
	InflationRate         float64 `json:"inflationRate"         form:"inflationRate"`
	InflationRateCheckbox string  `json:"inflationRateCheckbox" form:"inflationRateCheckbox"`
	TakeoutRate           float64 `json:"takeoutRate"           form:"takeoutRate"`
	TakeoutRateCheckbox   string  `json:"takeoutRateCheckbox"   form:"takeoutRateCheckbox"`
	Tax                   float64 `json:"tax"                   form:"tax"`
	TaxCheckbox           string  `json:"taxCheckbox"           form:"taxCheckbox"`
}

func (r *RawUserInputs) ToUserInputs() UserInputs {

	startCapital := NumberInputToInt(r.StartCapital)
	savingsRate := NumberInputToInt(r.SavingsRate)
	inflationRateCheckbox := CheckboxInputToBool(r.InflationRateCheckbox)
	takeoutRateCheckbox := CheckboxInputToBool(r.TakeoutRateCheckbox)
	taxCheckbox := CheckboxInputToBool(r.TaxCheckbox)

	return UserInputs{
		StartCapital:          startCapital,
		SavingsRate:           savingsRate,
		Years:                 r.Years,
		AnnualReturnRate:      r.AnnualReturnRate,
		InflationRate:         r.InflationRate,
		InflationRateCheckbox: inflationRateCheckbox,
		TakeoutRate:           r.TakeoutRate,
		TakeoutRateCheckbox:   takeoutRateCheckbox,
		Tax:                   r.Tax,
		TaxCheckbox:           taxCheckbox,
	}
}

type UserInputs struct {
	StartCapital          int
	SavingsRate           int
	Years                 int
	AnnualReturnRate      float64
	InflationRate         float64
	InflationRateCheckbox bool
	TakeoutRate           float64
	TakeoutRateCheckbox   bool
	Tax                   float64
	TaxCheckbox           bool
}
