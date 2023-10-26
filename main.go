package main

import (
	"os"

	"github.com/joel-beck/sparplanrechner/pkg/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":" + port))
}
