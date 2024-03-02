package handler

import (
	"github.com/joel-beck/sparplanrechner/pkg/calculator"
	"github.com/joel-beck/sparplanrechner/pkg/converters"
	"github.com/joel-beck/sparplanrechner/pkg/logging"
	"github.com/joel-beck/sparplanrechner/pkg/types"
	"github.com/labstack/echo/v4"
)

func bindUserInputs(c echo.Context) *types.UserInputs {
	userInputs := new(types.UserInputs)
	if err := c.Bind(userInputs); err != nil {
		c.Logger().Error(err)
	}
	logging.LogUserInputs(c, *userInputs)
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
