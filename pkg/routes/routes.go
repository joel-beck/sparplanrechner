package routes

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/joel-beck/sparplanrechner/pkg/calculator"
	"github.com/labstack/echo/v4"
)

var tmpl *template.Template

func InitRoutes(e *echo.Echo) {
	tmpl = template.Must(template.ParseFiles("web/index.html", "templates/result_template.html"))

	e.GET("/", ParseTemplates)
	e.Static("/", "web")
	e.POST("/calculate", sendResponse)
}

func ParseTemplates(c echo.Context) error {
	return tmpl.ExecuteTemplate(c.Response().Writer, "index.html", nil)
}

func collectTemplateData(amounts calculator.Amounts, startCapital int) map[string]interface{} {
	years := make([]int, len(amounts.AnnualTotals))
	for i := range years {
		years[i] = i + 1
	}

	return map[string]interface{}{
		"Years":          years,
		"Total":          calculator.FormatTotalAmount(amounts.AnnualTotals),
		"TotalPayments":  calculator.FormatTotalPayments(amounts.MonthlyPayments, startCapital),
		"TotalReturns":   calculator.FormatTotalReturns(amounts.MonthlyReturns),
		"AnnualTotals":   calculator.FormatAnnualTotals(amounts.AnnualTotals),
		"AnnualPayments": calculator.FormatAnnualPayments(amounts.MonthlyPayments, startCapital),
		"AnnualReturns":  calculator.FormatAnnualReturns(amounts.MonthlyReturns),
	}
}

func sendResponse(c echo.Context) error {
	req := new(calculator.InvestmentPlanRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input parameters"})
	}

	amounts := calculator.CalculateAmounts(req.StartCapital, req.SavingsRate, req.AnnualReturn, req.Years)
	data := collectTemplateData(amounts, req.StartCapital)

	t, err := template.ParseFiles("templates/result_template.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	result := tpl.String()
	return c.HTML(http.StatusOK, result)
}
