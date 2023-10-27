package routes

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/joel-beck/sparplanrechner/pkg/types"
	"github.com/labstack/echo/v4"
)

var tmpl *template.Template

func InitRoutes(e *echo.Echo) {
	tmpl = template.Must(template.ParseFiles("web/index.html", "templates/results.html"))

	e.GET("/", ParseTemplates)
	e.Static("/", "web")
	e.POST("/calculate", SendResponse)
}

func ParseTemplates(c echo.Context) error {
	return tmpl.ExecuteTemplate(c.Response().Writer, "index.html", nil)
}

// BindRequest binds the incoming request data to the UserInputs struct
func BindRequest(c echo.Context, req *types.UserInputs) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	return nil
}

// ParseAndExecuteTemplate parses the HTML template and executes it with the given data
func ParseAndExecuteTemplate(templateData map[string]interface{}) (string, error) {
	t, err := template.ParseFiles("templates/results.html")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, templateData); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

// SendResponse handles the response logic
func SendResponse(c echo.Context) error {
	inputs := types.UserInputs{}

	if err := BindRequest(c, &inputs); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input parameters"})
	}

	// Log User Inputs
	c.Logger().Infof("User Inputs: %+v", inputs)

	amounts := types.CalculateAmounts(inputs)

	monthlyIntermediateTotals := types.MonthlyIntermediateTotals{
		MonthlyPayments: amounts.MonthlyPayments,
		MonthlyReturns:  amounts.MonthlyReturns,
	}
	annualIntermediateTotals := types.MonthlyToAnnualTotals(
		amounts.AnnualTotals,
		amounts.InflationDiscountedAnnualTotals,
		monthlyIntermediateTotals,
		inputs.StartCapital,
	)

	totals := types.TotalsFromIntermediates(annualIntermediateTotals)
	takeouts := types.TakeoutsFromTotal(totals.Total, inputs)

	templateData := types.CollectTemplateData(annualIntermediateTotals, totals, takeouts, inputs.StartCapital)

	result, err := ParseAndExecuteTemplate(templateData)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Log Response Data
	c.Logger().Infof("Response Data: %+v", templateData)

	return c.HTML(http.StatusOK, result)
}
