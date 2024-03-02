package main

import (
	"os"

	"github.com/joel-beck/sparplanrechner/pkg/handler"
	"github.com/joel-beck/sparplanrechner/pkg/logging"
	"github.com/joel-beck/sparplanrechner/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func setPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func main() {
	logging.SetupLogger(zerolog.DebugLevel)

	e := echo.New()
	e.Use(middleware.ZerologMiddleware())

	// map all files from the local `dist` directory to the `/static` route access e.g.
	// via http://localhost:8080/static/style.css Note that the relative file path to
	// the `dist` directory always starts from the project root
	e.Static("/static", "./web/dist")

	e.GET("/", handler.HandleIndex)
	e.GET("/health", handler.HandleHealthCheck)
	e.POST("/calculate", handler.HandleFormResults)

	port := setPort()
	e.Logger.Fatal(e.Start(":" + port))
}
