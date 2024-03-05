package types

import (
	"strconv"
	"strings"

	"github.com/rs/zerolog"
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
	StartCapital          string  `json:"startCapitalInput"          form:"startCapitalInput"`
	SavingsRate           int     `json:"savingsRateInput"           form:"savingsRateInput"`
	Years                 int     `json:"yearsInput"                 form:"yearsInput"`
	ReturnRate            float64 `json:"returnRateInput"            form:"returnRateInput"`
	InflationRate         float64 `json:"inflationRateInput"         form:"inflationRateInput"`
	InflationRateCheckbox string  `json:"inflationRateCheckbox"      form:"inflationRateCheckbox"`
	TakeoutRate           float64 `json:"takeoutRateInput"           form:"takeoutRateInput"`
	TakeoutRateCheckbox   string  `json:"takeoutRateCheckbox"        form:"takeoutRateCheckbox"`
	Tax                   float64 `json:"taxInput"                   form:"taxInput"`
	TaxCheckbox           string  `json:"taxCheckbox"                form:"taxCheckbox"`
}

func (r RawUserInputs) ToUserInputs() UserInputs {

	startCapital := NumberInputToInt(r.StartCapital)
	inflationRateCheckbox := CheckboxInputToBool(r.InflationRateCheckbox)
	takeoutRateCheckbox := CheckboxInputToBool(r.TakeoutRateCheckbox)
	taxCheckbox := CheckboxInputToBool(r.TaxCheckbox)

	return UserInputs{
		StartCapital:          startCapital,
		SavingsRate:           r.SavingsRate,
		Years:                 r.Years,
		ReturnRate:            r.ReturnRate,
		InflationRate:         r.InflationRate,
		InflationRateCheckbox: inflationRateCheckbox,
		TakeoutRate:           r.TakeoutRate,
		TakeoutRateCheckbox:   takeoutRateCheckbox,
		Tax:                   r.Tax,
		TaxCheckbox:           taxCheckbox,
	}
}

func (r RawUserInputs) MarshalZerologObject(e *zerolog.Event) {
	e.Str("startCapital", r.StartCapital).
		Int("savingsRate", r.SavingsRate).
		Int("years", r.Years).
		Float64("returnRate", r.ReturnRate).
		Float64("inflationRate", r.InflationRate).
		Str("inflationRateCheckbox", r.InflationRateCheckbox).
		Float64("takeoutRate", r.TakeoutRate).
		Str("takeoutRateCheckbox", r.TakeoutRateCheckbox).
		Float64("tax", r.Tax).
		Str("taxCheckbox", r.TaxCheckbox)

}

type UserInputs struct {
	StartCapital          int
	SavingsRate           int
	Years                 int
	ReturnRate            float64
	InflationRate         float64
	InflationRateCheckbox bool
	TakeoutRate           float64
	TakeoutRateCheckbox   bool
	Tax                   float64
	TaxCheckbox           bool
}

func (u UserInputs) MarshalZerologObject(e *zerolog.Event) {
	e.Int("startCapital", u.StartCapital).
		Int("savingsRate", u.SavingsRate).
		Int("years", u.Years).
		Float64("returnRate", u.ReturnRate).
		Float64("inflationRate", u.InflationRate).
		Bool("inflationRateCheckbox", u.InflationRateCheckbox).
		Float64("takeoutRate", u.TakeoutRate).
		Bool("takeoutRateCheckbox", u.TakeoutRateCheckbox).
		Float64("tax", u.Tax).
		Bool("taxCheckbox", u.TaxCheckbox)
}
