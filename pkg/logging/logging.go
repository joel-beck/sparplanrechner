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

	fmt.Printf("startCapital: %s\n", c.FormValue("startCapital"))
	fmt.Printf("savingsRate: %s\n", c.FormValue("savingsRate"))
	fmt.Printf("annualReturn: %s\n", c.FormValue("annualReturn"))
	fmt.Printf("years: %s\n", c.FormValue("years"))
	fmt.Printf("inflationRate: %s\n", c.FormValue("inflationRate"))

	fmt.Println()
}

func LogRequest(req *calculator.UserInputs) {
	requestLog, _ := json.MarshalIndent(req, "", "  ")
	fmt.Printf("Received request: %s\n", requestLog)
	fmt.Println()
}
