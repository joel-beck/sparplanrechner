package main

import (
	"github.com/joel-beck/sparplanrechner/pkg/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", routes.ParseTemplates)
	e.Static("/", "web")
	e.POST("/calculate", routes.ProcessUserInputs)

	e.Logger.Fatal(e.Start(":8080"))
}
