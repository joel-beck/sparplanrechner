package main

import (
	"html/template"
	"os"

	"github.com/joel-beck/sparplanrechner/pkg/logging"
	"github.com/joel-beck/sparplanrechner/pkg/middleware"
	"github.com/joel-beck/sparplanrechner/pkg/routes"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func parseTemplates() *template.Template {

	// `ParseGlob` does not parse all templates in nested directories, i.e.
	// `ParseGlob("web/**/*.html")` does not parse the `web/index.html` template!
	parsedTemplates := template.Must(template.ParseGlob("web/*.html"))
	template.Must(parsedTemplates.ParseGlob("web/views/*.html"))
	return parsedTemplates
}

func setPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func main() {
	logging.SetupLogger(zerolog.DebugLevel)

	parsedTemplates := parseTemplates()
	logging.LogParsedTemplates(parsedTemplates)

	e := echo.New()
	e.Use(middleware.ZerologMiddleware())

	e.Renderer = &routes.TemplateRenderer{
		Templates: parsedTemplates,
	}

	handler := &routes.Handler{
		Renderer: e.Renderer.(*routes.TemplateRenderer),
	}

	e.GET("/", handler.Index)

	// serve static files such as CSS and JS
	e.Static("/", "web")

	e.POST("/calculate", handler.Calculate)

	port := setPort()
	e.Logger.Fatal(e.Start(":" + port))
}
