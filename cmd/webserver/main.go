package main

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"

	"github.com/joel-beck/sparplanrechner/pkg/investmentcalc"
	"github.com/labstack/echo/v4"
)

func sendHtmlResponse(c echo.Context, req *investmentcalc.InvestmentPlanRequest) error {
	amounts := investmentcalc.CalculateAmounts(req.InitialCapital, req.AnnualReturn, req.Years)

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

func calculateInvestmentPlan(c echo.Context) error {
	initialCapital, err := strconv.ParseFloat(c.FormValue("initialCapital"), 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid initial capital")
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
		AnnualReturn:   annualReturn,
		Years:          years,
	}

	return sendHtmlResponse(c, req)
}

func main() {
	e := echo.New()

	funcMap := template.FuncMap{
		"even": func(i int) bool {
			return i%2 == 0
		},
	}

	tmpl, err := template.New("").Funcs(funcMap).ParseFiles("web/index.html", "templates/result_template.html")
	if err != nil {
		panic("Failed to parse HTML templates")
	}

	e.GET("/", func(c echo.Context) error {
		return tmpl.ExecuteTemplate(c.Response().Writer, "index.html", nil)
	})

	e.Static("/", "web")
	e.POST("/calculate", calculateInvestmentPlan)
	e.Logger.Fatal(e.Start(":8080"))
}
