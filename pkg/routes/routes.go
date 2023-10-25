package routes

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"

	"github.com/joel-beck/sparplanrechner/pkg/calculator"
	"github.com/labstack/echo/v4"
)

func ParseTemplates(c echo.Context) error {
	tmpl, err := template.ParseFiles("web/index.html", "templates/result_template.html")
	if err != nil {
		panic("Failed to parse HTML templates")
	}

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

func sendHtmlResponse(c echo.Context, req *calculator.InvestmentPlanRequest) error {
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

func ProcessUserInputs(c echo.Context) error {
	startCapital, err := strconv.Atoi(c.FormValue("startCapital"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid starting capital")
	}

	savingsRate, err := strconv.Atoi(c.FormValue("savingsRate"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid savings rate")
	}

	annualReturn, err := strconv.ParseFloat(c.FormValue("annualReturn"), 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid annual return")
	}

	years, err := strconv.Atoi(c.FormValue("years"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid number of years")
	}

	req := &calculator.InvestmentPlanRequest{
		StartCapital: startCapital,
		SavingsRate:  savingsRate,
		AnnualReturn: annualReturn,
		Years:        years,
	}

	return sendHtmlResponse(c, req)
}
