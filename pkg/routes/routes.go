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
	e.POST("/calculate", SendResponse)
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
		"Years":                           years,
		"Total":                           calculator.FormatTotalAmount(amounts.AnnualTotals),
		"TotalInflationDiscounted":        calculator.FormatTotalAmount(amounts.InflationDiscountedAnnualTotals),
		"TotalPayments":                   calculator.FormatTotalPayments(amounts.MonthlyPayments, startCapital),
		"TotalReturns":                    calculator.FormatTotalReturns(amounts.MonthlyReturns),
		"AnnualTotals":                    calculator.FormatAnnualTotals(amounts.AnnualTotals),
		"AnnualInflationDiscountedTotals": calculator.FormatAnnualTotals(amounts.InflationDiscountedAnnualTotals),
		"AnnualPayments":                  calculator.FormatAnnualPayments(amounts.MonthlyPayments, startCapital),
		"AnnualReturns":                   calculator.FormatAnnualReturns(amounts.MonthlyReturns),
	}
}

// BindRequest binds the incoming request data to the InvestmentPlanRequest struct
func BindRequest(c echo.Context, req *calculator.UserInputs) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	return nil
}

// ParseAndExecuteTemplate parses the HTML template and executes it with the given data
func ParseAndExecuteTemplate(data map[string]interface{}) (string, error) {
	t, err := template.ParseFiles("templates/result_template.html")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

// SendResponse handles the response logic
func SendResponse(c echo.Context) error {
	req := new(calculator.UserInputs)

	if err := BindRequest(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input parameters"})
	}

	// Log User Inputs
	c.Logger().Infof("User Inputs: %+v", req)

	amounts := calculator.CalculateAmounts(
		req.StartCapital,
		req.SavingsRate,
		req.AnnualReturn,
		req.Years,
		req.InflationRate,
	)
	data := collectTemplateData(amounts, req.StartCapital)

	result, err := ParseAndExecuteTemplate(data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Log Response Data
	c.Logger().Infof("Response Data: %+v", data)

	return c.HTML(http.StatusOK, result)
}
