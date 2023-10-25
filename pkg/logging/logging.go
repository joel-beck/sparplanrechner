package logging

import (
	"encoding/json"
	"fmt"

	"github.com/joel-beck/sparplanrechner/pkg/calculator"
	"github.com/labstack/echo/v4"
)

func LogUserInputs(c echo.Context) {
	fmt.Println()

	fmt.Println("Received user inputs:")
	fmt.Println("----------------------")

	fmt.Printf("initialCapital: %s\n", c.FormValue("initialCapital"))
	fmt.Printf("annualReturn: %s\n", c.FormValue("annualReturn"))
	fmt.Printf("years: %s\n", c.FormValue("years"))

	fmt.Println()
}

func LogRequest(req *calculator.InvestmentPlanRequest) {
	requestLog, _ := json.MarshalIndent(req, "", "  ")
	fmt.Printf("Received request: %s\n", requestLog)
	fmt.Println()
}
