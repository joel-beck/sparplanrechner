package handler

import (
	"github.com/joel-beck/sparplanrechner/pkg/templates"
	"github.com/joel-beck/sparplanrechner/pkg/templates/layouts"
	"github.com/joel-beck/sparplanrechner/pkg/templates/results"
	"github.com/labstack/echo/v4"
)

func HandleIndex(c echo.Context) error {
	return templates.Render(c, layouts.Index())
}

func HandleHealthCheck(c echo.Context) error {
	return c.String(200, "OK")
}

func HandleFormResults(c echo.Context) error {
	userInputs := bindUserInputs(c)
	responseData := computeResponseData(c, *userInputs)

	return templates.Render(c, results.FormResults(responseData))
}
