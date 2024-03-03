package handler

import (
	"net/http"

	"github.com/joel-beck/sparplanrechner/internal/templates"
	"github.com/joel-beck/sparplanrechner/internal/templates/layouts"
	"github.com/joel-beck/sparplanrechner/internal/templates/results"
	"github.com/labstack/echo/v4"
)

func HandleIndex(c echo.Context) error {
	return templates.Render(c, layouts.Index())
}

func HandleHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func HandleFormResults(c echo.Context) error {
	rawUserInputs := bindRawUserInputs(c)
	userInputs := parseUserInputs(c, rawUserInputs)
	responseData := computeResponseData(c, userInputs)

	return templates.Render(c, results.FormResults(responseData))
}
