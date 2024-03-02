package handler

import (
	"github.com/joel-beck/sparplanrechner/internal/calculator"
	"github.com/joel-beck/sparplanrechner/internal/converters"
	"github.com/joel-beck/sparplanrechner/internal/logging"
	"github.com/joel-beck/sparplanrechner/internal/types"
	"github.com/labstack/echo/v4"
)

func bindRawUserInputs(c echo.Context) *types.RawUserInputs {
	rawUserInputs := new(types.RawUserInputs)
	if err := c.Bind(rawUserInputs); err != nil {
		c.Logger().Error(err)
	}
	logging.LogRawUserInputs(c, *rawUserInputs)
	return rawUserInputs
}

func parseUserInputs(c echo.Context, rawUserInputs *types.RawUserInputs) types.UserInputs {
	userInputs := rawUserInputs.ToUserInputs()
	logging.LogUserInputs(c, userInputs)
	return userInputs
}

func computeResponseData(c echo.Context, userInputs types.UserInputs) types.ResponseData {
	annualIntermediateAmounts := calculator.ComputeAnnualTotals(userInputs)
	totals := converters.TotalsFromIntermediates(annualIntermediateAmounts)
	takeouts := converters.TakeoutsFromTotal(totals.Total, userInputs)

	responseData := types.CollectResponseData(annualIntermediateAmounts, totals, takeouts, userInputs.StartCapital)

	logging.LogResponseData(c, responseData)
	return responseData
}
