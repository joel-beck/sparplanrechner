package routes

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"

	"github.com/joel-beck/sparplanrechner/pkg/calculator"
	"github.com/joel-beck/sparplanrechner/pkg/converters"
	"github.com/joel-beck/sparplanrechner/pkg/middleware"
	"github.com/joel-beck/sparplanrechner/pkg/types"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var tmpl *template.Template

func logUserInputs(c echo.Context, inputs types.UserInputs) {
	log.Debug().
		Interface("user_inputs", inputs).
		Msg("User Inputs")
}

func logResponse(c echo.Context, templateData map[string]interface{}) {
	log.Debug().
		Interface("response_data", templateData).
		Msg("Response Data")
}

// BindRequest binds the incoming request data to the UserInputs struct
func BindRequest(c echo.Context, req *types.UserInputs) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	return nil
}

// TODO: Clean up naming of parse and execute template functions
func executeTemplate(c echo.Context) error {
	return tmpl.ExecuteTemplate(c.Response().Writer, "index.html", nil)
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

func parseTemplate(
	c echo.Context,
	inputs types.UserInputs,
	annualIntermediateTotals types.AnnualIntermediateTotals,
	totals types.Totals,
	takeouts types.Takeouts,
) (string, error) {
	templateData := types.CollectTemplateData(annualIntermediateTotals, totals, takeouts, inputs.StartCapital)

	result, err := ParseAndExecuteTemplate(templateData)
	if err != nil {
		return "", c.String(http.StatusInternalServerError, err.Error())
	}

	logResponse(c, templateData)
	return result, nil
}

func HTMLResponse(c echo.Context, result string) error {
	return c.HTML(http.StatusOK, result)
}

// SendResponse handles the response logic
func SendResponse(c echo.Context) error {
	inputs := types.UserInputs{}

	if err := BindRequest(c, &inputs); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input parameters"})
	}
	logUserInputs(c, inputs)

	annualIntermediateTotals := calculator.ComputeAnnualTotals(inputs)
	totals := converters.TotalsFromIntermediates(annualIntermediateTotals)
	takeouts := converters.TakeoutsFromTotal(totals.Total, inputs)

	result, err := parseTemplate(c, inputs, annualIntermediateTotals, totals, takeouts)
	if err != nil {
		return errors.New("Error while parsing template")
	}

	return HTMLResponse(c, result)
}

func InitRoutes(e *echo.Echo) {
	tmpl = template.Must(template.ParseFiles("web/index.html", "templates/results.html"))

	e.Use(middleware.ZerologMiddleware())

	e.GET("/", executeTemplate)
	e.Static("/", "web")
	e.POST("/calculate", SendResponse)
}
