package routes

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"

	"github.com/joel-beck/sparplanrechner/pkg/investmentcalc"
	"github.com/labstack/echo/v4"
)

func ParseTemplates(c echo.Context) error {
	tmpl, err := template.ParseFiles("web/index.html", "templates/result_template.html")
	if err != nil {
		panic("Failed to parse HTML templates")
	}

	return tmpl.ExecuteTemplate(c.Response().Writer, "index.html", nil)
}

func sendHtmlResponse(c echo.Context, req *investmentcalc.InvestmentPlanRequest) error {
	amounts := investmentcalc.CalculateAmounts(req.InitialCapital, req.SavingsRate, req.AnnualReturn, req.Years)

	finalAmount := investmentcalc.GetFinalAmount(amounts)
	intermediateAmounts := investmentcalc.GetIntermediateAmounts(amounts)

	t, err := template.ParseFiles("templates/result_template.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	data := map[string]interface{}{
		"FinalAmount":         finalAmount,
		"IntermediateAmounts": intermediateAmounts,
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	result := tpl.String()
	return c.HTML(http.StatusOK, result)
}

func CalculateInvestmentPlan(c echo.Context) error {
	initialCapital, err := strconv.Atoi(c.FormValue("initialCapital"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid initial capital")
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

	req := &investmentcalc.InvestmentPlanRequest{
		InitialCapital: initialCapital,
		SavingsRate:    savingsRate,
		AnnualReturn:   annualReturn,
		Years:          years,
	}

	return sendHtmlResponse(c, req)
}
