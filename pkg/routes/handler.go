package routes

import (
	"net/http"

	"github.com/joel-beck/sparplanrechner/pkg/calculator"
	"github.com/joel-beck/sparplanrechner/pkg/converters"
	"github.com/joel-beck/sparplanrechner/pkg/logging"
	"github.com/joel-beck/sparplanrechner/pkg/types"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Renderer *TemplateRenderer
}

func (h *Handler) Index(c echo.Context) error {
	return h.Renderer.Render(c.Response().Writer, "index", nil, c)
}

func (h *Handler) Calculate(c echo.Context) error {
	inputs := types.UserInputs{}
	logging.LogUserInputs(c, inputs)

	if err := c.Bind(&inputs); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input parameters"})
	}

	responseData := retrieveResponseData(inputs)
	logging.LogResponseData(c, responseData)

	return h.Renderer.Render(c.Response().Writer, "htmxResults", responseData, c)
}

func retrieveResponseData(userInputs types.UserInputs) types.ResponseData {
	annualIntermediateTotals := calculator.ComputeAnnualTotals(userInputs)
	totals := converters.TotalsFromIntermediates(annualIntermediateTotals)
	takeouts := converters.TakeoutsFromTotal(totals.Total, userInputs)

	return types.CollectResponseData(annualIntermediateTotals, totals, takeouts, userInputs.StartCapital)
}
